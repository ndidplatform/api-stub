'use strict';


/**
 * Get Data requested from IDP and AS
 * TBD
 *
 * request_id String Request ID
 * returns List
 **/
exports.get_request_data = function(request_id) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = [ {
  "source_node_id" : "source_node_id",
  "data" : "data",
  "service_id" : "service_id",
  "source_signature" : "source_signature"
}, {
  "source_node_id" : "source_node_id",
  "data" : "data",
  "service_id" : "service_id",
  "source_signature" : "source_signature"
} ];
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * get request_id for the request with {reference_number} -- Only request from this RP will succeed, reference number is not propagated to other nodes.
 *
 * reference_number String reference number
 * returns String
 **/
exports.get_request_id_from_reference_number = function(reference_number) {
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
 * Fetch Request status
 * TBD
 *
 * request_id String request ID
 * returns RequestsStatus
 **/
exports.get_request_status = function(request_id) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "min_aal" : 6,
  "request_type" : "consent",
  "request_message" : "request_message",
  "min_idp" : 1,
  "service_id_list" : [ "service_id_list", "service_id_list" ],
  "min_ial" : 0,
  "timeout" : 5,
  "status" : "open"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Send request to {namespace}/{identifier}
 * asdf
 *
 * namespace String ID Namespace
 * identifier String Identifier for the ID
 * request Requests Additional identifiers for this ID to add
 * timeout List timeout in ms.  Once this timeout is reached, the API returns control to the caller, and caller should poll for update.  If empty, assume indefinite sync mode. (not safe) (optional)
 * returns String
 **/
exports.send_request_to_id = function(namespace,identifier,request,timeout) {
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

