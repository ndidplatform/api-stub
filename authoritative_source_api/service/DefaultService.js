'use strict';


/**
 * Get information about service offered by this AS
 * asdf
 *
 * service_id String service ID
 * returns ASServices
 **/
exports.get_service = function(service_id) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "min_aal" : 6,
  "service_name" : "service_name",
  "service_id" : "service_id",
  "min_ial" : 0,
  "url" : "url"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Add or Update service offered by this AS
 * asdf
 *
 * service_id String service ID
 * service ASServices IDP_response
 * no response value expected for this operation
 **/
exports.register_service = function(service_id,service) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}

