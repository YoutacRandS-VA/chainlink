// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import {RouterBase} from "./RouterBase.sol";
import {IFunctionsRouter} from "./interfaces/IFunctionsRouter.sol";
import {IVersioned} from "./interfaces/IVersioned.sol";
import {IFunctionsCoordinator} from "./interfaces/IFunctionsCoordinator.sol";
import {IFunctionsBilling} from "./interfaces/IFunctionsBilling.sol";
import {AuthorizedOriginReceiver} from "./accessControl/AuthorizedOriginReceiver.sol";
import {FunctionsSubscriptions} from "./FunctionsSubscriptions.sol";

contract FunctionsRouter is RouterBase, IFunctionsRouter, AuthorizedOriginReceiver, FunctionsSubscriptions {
  // ================================================================
  // |                    Configuration state                       |
  // ================================================================
  struct Config {
    // Flat fee (in Juels of LINK) that will be paid to the Router owner for operation of the network
    uint96 adminFee;
  }
  Config private s_config;
  event ConfigSet(uint96 adminFee);

  error OnlyCallableByRoute();

  // ================================================================
  // |                       Initialization                         |
  // ================================================================
  constructor(
    uint16 timelockBlocks,
    uint16 maximumTimelockBlocks,
    bool useAllowList,
    address linkToken,
    string[] memory initialLabels,
    address[] memory initialAddresses,
    bytes memory config
  )
    RouterBase(
      "FunctionsRouter",
      msg.sender,
      timelockBlocks,
      maximumTimelockBlocks,
      initialLabels,
      initialAddresses,
      config
    )
    AuthorizedOriginReceiver(useAllowList)
    FunctionsSubscriptions(linkToken)
  {}

  // ================================================================
  // |                 Configuration methods                        |
  // ================================================================
  /**
   * @notice Sets the configuration for FunctionsRouter specific state
   * @param config bytes of config data to set the following:
   *  - adminFee: fee that will be paid to the Router owner for operating the network
   */
  function _setConfig(bytes memory config) internal override {
    uint96 adminFee = abi.decode(config, (uint96));
    s_config = Config({adminFee: adminFee});
    emit ConfigSet(adminFee);
  }

  /**
   * @inheritdoc IFunctionsRouter
   */
  function getAdminFee() external view override returns (uint96) {
    return s_config.adminFee;
  }

  // ================================================================
  // |                      Request methods                         |
  // ================================================================

  function _smoke(bytes calldata data) internal override onlyAuthorizedUsers returns (bytes32) {
    address route = this.getRoute("FunctionsCoordinator", true);
    IFunctionsCoordinator coordinator = IFunctionsCoordinator(route);
    return coordinator.sendRequest(msg.sender, data);
  }

  /**
   * @inheritdoc IFunctionsRouter
   */
  function sendRequest(bytes calldata data) external override onlyAuthorizedUsers returns (bytes32) {
    address route = this.getRoute("FunctionsCoordinator", false);
    IFunctionsCoordinator coordinator = IFunctionsCoordinator(route);
    return coordinator.sendRequest(msg.sender, data);
  }

  /**
   * @inheritdoc IFunctionsRouter
   */
  function timeoutRequests(bytes32[] calldata requestIdsToTimeout) external onlyAuthorizedUsers {
    address route = this.getRoute("FunctionsCoordinator", false);
    IFunctionsBilling coordinator = IFunctionsBilling(route);
    coordinator.timeoutRequests(requestIdsToTimeout);
  }

  // ================================================================
  // |                    Modifier Overrides                        |
  // ================================================================

  modifier onlyRoute() override {
    // Since Proposal Sets are not deleted, we can check the:
    // - "current" and "next" version if the Proposal Set has not been applied
    // - "current" and "previous" version if the Proposal Set has been applied
    // This allows in-flight requests to still complete after an upgrade has taken effect
    bool isRoute = false;
    for (uint8 i = 0; i < s_proposalSet.labels.length; i++) {
      if (msg.sender == s_proposalSet.from[i] || msg.sender == s_proposalSet.to[i]) {
        isRoute = true;
        break;
      }
    }
    if (isRoute == false) {
      revert OnlyCallableByRoute();
    }
    _;
  }

  function _canSetAuthorizedSenders() internal view override onlyOwner returns (bool) {
    return msg.sender == owner();
  }

  modifier onlyAuthorizedUsers() override {
    _validateIsAuthorizedSender();
    _;
  }

  modifier nonReentrant() override {
    address route = this.getRoute("FunctionsCoordinator", true);
    IFunctionsBilling coordinator = IFunctionsBilling(route);
    if (coordinator.isReentrancyLocked()) {
      revert Reentrant();
    }
    _;
  }
  modifier onlyRouterOwner() override {
    _validateOwnership();
    _;
  }
}
