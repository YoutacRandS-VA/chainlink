// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import {IFunctionsCoordinator} from "./interfaces/IFunctionsCoordinator.sol";
import {IFunctionsSubscriptions} from "./interfaces/IFunctionsSubscriptions.sol";
import {OCR2Base} from "./ocr/OCR2Base.sol";
import {IFunctionsBilling, FunctionsBilling, Route} from "./FunctionsBilling.sol";
import {Functions} from "./Functions.sol";
import {IOwnable} from "../../../shared/interfaces/IOwnable.sol";

/**
 * @title Functions Coordinator contract
 * @notice Contract that nodes of a Decentralized Oracle Network (DON) interact with
 * @dev THIS CONTRACT HAS NOT GONE THROUGH ANY SECURITY REVIEW. DO NOT USE IN PROD.
 */
contract FunctionsCoordinator is OCR2Base, IFunctionsCoordinator, FunctionsBilling {
  uint16 public constant REQUEST_DATA_VERSION = Functions.REQUEST_DATA_VERSION;

  event OracleRequest(
    bytes32 indexed requestId,
    address requestingContract,
    address requestInitiator,
    uint64 subscriptionId,
    address subscriptionOwner,
    bytes data
  );
  event OracleResponse(bytes32 indexed requestId);
  event UserCallbackError(bytes32 indexed requestId, string reason);
  event InvalidRequestID(bytes32 indexed requestId);

  error UnsupportedRequestDataVersion();
  error EmptyRequestData();
  error InconsistentReportData();
  error EmptyPublicKey();
  error UnauthorizedPublicKeyChange();

  bytes private s_donPublicKey;
  mapping(address => bytes) private s_nodePublicKeys;

  constructor(
    string memory id,
    uint16 version,
    address router,
    bytes memory config,
    address linkToNativeFeed
  ) OCR2Base(true) FunctionsBilling(id, version, router, config, linkToNativeFeed) {}

  /**
   * @inheritdoc IFunctionsCoordinator
   */
  function getDONPublicKey() external view override returns (bytes memory) {
    return s_donPublicKey;
  }

  /**
   * @inheritdoc IFunctionsCoordinator
   */
  function setDONPublicKey(bytes calldata donPublicKey) external override onlyRouterOwner {
    if (donPublicKey.length == 0) {
      revert EmptyPublicKey();
    }
    s_donPublicKey = donPublicKey;
  }

  /**
   * @dev check if node is in current transmitter list
   */
  function _isTransmitter(address node) internal view returns (bool) {
    address[] memory nodes = this.transmitters();
    for (uint256 i = 0; i < nodes.length; i++) {
      if (nodes[i] == node) {
        return true;
      }
    }
    return false;
  }

  /**
   * @inheritdoc IFunctionsCoordinator
   */
  function setNodePublicKey(address node, bytes calldata publicKey) external override {
    // Owner can set anything. Transmitters can set only their own key.
    if (!(msg.sender == IOwnable(address(s_router)).owner() || (_isTransmitter(msg.sender) && msg.sender == node))) {
      revert UnauthorizedPublicKeyChange();
    }
    s_nodePublicKeys[node] = publicKey;
  }

  /**
   * @inheritdoc IFunctionsCoordinator
   */
  function deleteNodePublicKey(address node) external override {
    // Owner can delete anything. Others can delete only their own key.
    if (!(msg.sender == IOwnable(address(s_router)).owner() || msg.sender == node)) {
      revert UnauthorizedPublicKeyChange();
    }
    delete s_nodePublicKeys[node];
  }

  /**
   * @inheritdoc IFunctionsCoordinator
   */
  function getAllNodePublicKeys() external view override returns (address[] memory, bytes[] memory) {
    address[] memory nodes = this.transmitters();
    bytes[] memory keys = new bytes[](nodes.length);
    for (uint256 i = 0; i < nodes.length; i++) {
      keys[i] = s_nodePublicKeys[nodes[i]];
    }
    return (nodes, keys);
  }

  /**
   * @inheritdoc IFunctionsCoordinator
   */
  function sendRequest(address caller, bytes calldata requestData) external override onlyRouter returns (bytes32) {
    if (requestData.length == 0) {
      revert EmptyRequestData();
    }

    (uint16 version, uint64 subscriptionId, uint32 gasLimit, bytes memory requestCBOR) = Functions.decodeRequest(
      requestData
    );

    if (version != REQUEST_DATA_VERSION) {
      revert UnsupportedRequestDataVersion();
    }

    IFunctionsSubscriptions subscriptions = IFunctionsSubscriptions(address(s_router));
    (, , address subscriptionOwner, ) = subscriptions.getSubscription(subscriptionId);
    if (subscriptionOwner == address(0)) {
      // TODO: gas optimize?
      revert InvalidSubscription();
    }

    // It's important to ensure that the consumer is in fact who they say they
    // are, otherwise they could use someone else's subscription balance.
    (bool allowed, , ) = subscriptions.getConsumer(caller, subscriptionId);

    if (allowed == false) {
      revert InvalidConsumer(subscriptionId, caller);
    }

    bytes32 requestId = startBilling(
      requestCBOR,
      IFunctionsBilling.RequestBilling(subscriptionId, caller, gasLimit, tx.gasprice)
    );

    emit OracleRequest(requestId, caller, tx.origin, subscriptionId, subscriptionOwner, requestCBOR);

    return requestId;
  }

  function _beforeSetConfig(uint8 _f, bytes memory _onchainConfig) internal override {}

  function _afterSetConfig(uint8 _f, bytes memory _onchainConfig) internal override {}

  function _validateReport(
    bytes32, /* configDigest */
    uint40, /* epochAndRound */
    bytes memory /* report */
  ) internal pure override returns (bool) {
    // validate within _report to save gas
    return true;
  }

  function _report(
    uint256 initialGas,
    address transmitter,
    uint8 signerCount,
    address[maxNumOracles] memory signers,
    bytes calldata report
  ) internal override {
    bytes32[] memory requestIds;
    bytes[] memory results;
    bytes[] memory errors;
    (
      requestIds,
      results,
      errors
      /*metadata,*/
    ) = abi.decode(report, (bytes32[], bytes[], bytes[])); // TODO: metadata through report
    if (requestIds.length == 0 || requestIds.length != results.length || requestIds.length != errors.length) {
      revert ReportInvalid();
    }

    uint256 reportValidationGasShare = (initialGas - gasleft()) / requestIds.length;

    for (uint256 i = 0; i < requestIds.length; i++) {
      IFunctionsBilling.FulfillResult result = fulfillAndBill(
        requestIds[i],
        results[i],
        errors[i],
        /* metadata[i], */
        transmitter,
        signers,
        signerCount,
        reportValidationGasShare,
        gasleft()
      );
      if (result == IFunctionsBilling.FulfillResult.USER_SUCCESS) {
        emit OracleResponse(requestIds[i]);
      } else if (result == IFunctionsBilling.FulfillResult.USER_ERROR) {
        emit UserCallbackError(requestIds[i], "error in callback");
      } else if (result == IFunctionsBilling.FulfillResult.INVALID_REQUEST_ID) {
        emit InvalidRequestID(requestIds[i]);
      }
    }
  }

  modifier onlyRouterOwner() override(OCR2Base, Route) {
    if (msg.sender != IOwnable(address(s_router)).owner()) {
      revert OnlyCallableByRouterOwner();
    }
    _;
  }
}
