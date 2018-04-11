'use strict';


/**
 * create node in the dpki module
 * NDID ONLY - used by NDID to add new organization to the network.  When an organization register to host NDID node, it must present the node_data to NDID (node_id, node_name, node_key, node_master_key)  node_id must be unique.
 *
 * node_data Node_data 
 * no response value expected for this operation
 **/
exports.dpkiNodeCreatePOST = function(node_data) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}


/**
 * Register a callback location for NDID node to call.
 * Register a callback location for NDID DPKI node to call to get things signed.  Required if want to use HSM or other methods of key signing and storage other than default (store at node app data)
 *
 * url String 
 * no response value expected for this operation
 **/
exports.dpkiNodeRegister_callbackPOST = function(url) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}


/**
 * Register a callback location for NDID node to call when it needs master key.
 * Register a callback location for NDID DPKI node to call to get things signed by master key.  Required if want to use HSM or other methods of key signing and storage other than default (store at node app data)
 *
 * url String 
 * no response value expected for this operation
 **/
exports.dpkiNodeRegister_callback_masterPOST = function(url) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}


/**
 * update node name
 * Update node_name, node_key or node_master_key
 *
 * node_data Node_data 
 * no response value expected for this operation
 **/
exports.dpkiNodeUpdatePOST = function(node_data) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}

