'use strict';


/**
 * Implement this function to sign using node's private key. 
 * Implement this function to sign using node's private key. (Must match the node's public key registered to the NDID node)
 *
 * sign_request Sign_request 
 * returns String
 **/
exports.dpkiMasterSignPOST = function(sign_request) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = "";
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Implement this function to sign using node's private key. 
 * Implement this function to sign using node's private key. (Must match the node's public key registered to the NDID node)
 *
 * sign_request Sign_request 
 * returns String
 **/
exports.dpkiSignPOST = function(sign_request) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = "";
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}

