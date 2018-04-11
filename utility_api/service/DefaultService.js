'use strict';


/**
 * Retrieve list of id of all available as given a service id
 * asdf
 *
 * service_id String Service ID
 * returns List
 **/
exports.get_as_id_list = function(service_id) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = [ {
  "min_aal" : 6,
  "as_id" : "as_id",
  "min_ial" : 0,
  "as_name" : "as_name"
}, {
  "min_aal" : 6,
  "as_id" : "as_id",
  "min_ial" : 0,
  "as_name" : "as_name"
} ];
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Retrieve all idp id
 * asdf
 *
 * min_ial Integer minimum ial suppored by IDP looking for (optional)
 * min_aal Integer minimum aal suppored by IDP looking for (optional)
 * returns List
 **/
exports.get_idp_list = function(min_ial,min_aal) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = [ {
  "min_aal" : 6,
  "idp_name" : "idp_name",
  "idp_id" : "idp_id",
  "min_ial" : 0
}, {
  "min_aal" : 6,
  "idp_name" : "idp_name",
  "idp_id" : "idp_id",
  "min_ial" : 0
} ];
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Retrieve all idp node_id relevant to/authorized by this {namespace}/{identifier}
 * asdf
 *
 * namespace String ID Namespace
 * identifier String Identifier for the ID
 * min_ial Integer minimum ial suppored by IDP looking for (optional)
 * min_aal Integer minimum aal suppored by IDP looking for (optional)
 * returns List
 **/
exports.get_idp_list_for_identifier = function(namespace,identifier,min_ial,min_aal) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = [ {
  "min_aal" : 6,
  "idp_name" : "idp_name",
  "idp_id" : "idp_id",
  "min_ial" : 0
}, {
  "min_aal" : 6,
  "idp_name" : "idp_name",
  "idp_id" : "idp_id",
  "min_ial" : 0
} ];
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}

