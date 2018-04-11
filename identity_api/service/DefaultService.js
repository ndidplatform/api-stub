'use strict';


/**
 * Add accessor method
 * asdf
 *
 * namespace String ID Namespace
 * identifier String Identifier for the ID
 * accessors Accessors endorsement
 * no response value expected for this operation
 **/
exports.add_accessor_method = function(namespace,identifier,accessors) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}


/**
 * Submit new identifiers for {namespace}/{identifier}
 * asdf
 *
 * namespace String ID Namespace
 * identifier String Identifier for the ID
 * new_identifiers Identifiers Additional identifiers for this ID to add
 * no response value expected for this operation
 **/
exports.add_identifier = function(namespace,identifier,new_identifiers) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}


/**
 * Fetch namespace/identifier
 * asdf
 *
 * namespace String ID Namespace
 * identifier String Identifier for the ID
 * no response value expected for this operation
 **/
exports.check_identity = function(namespace,identifier) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}


/**
 * Create Identity Record
 * IDP calls this method to create a new identity
 *
 * identity Identity Identity data needed to create identity.  Please note that 'namespace', 'identifier', and 'secret' are not stored directly.  Rather, Hash(namespace, identifier, secret) are created and stored.  Secret is not shared with any other IDP, rather is used by each IDP to prove it knows this person after having been granted permission to access this Identity.  In addition, this method provide privacy across IDP, as each IDP does not know what other IDP knows about this same person if IDP does not know the secret.
 * no response value expected for this operation
 **/
exports.create_identity = function(identity) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}


/**
 * Submit endorsement for {namespace}/{identifier}
 * asdf
 *
 * namespace String ID Namespace
 * identifier String Identifier for the ID
 * endorsements Identity endorsement
 * no response value expected for this operation
 **/
exports.endorse_identity = function(namespace,identifier,endorsements) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}


/**
 * Fetch Endorsement list for {namespace}/{identifier}
 * TBD
 *
 * namespace String ID Namespace
 * identifier String Identifier for the ID
 * returns Endorsements
 **/
exports.get_endorsement = function(namespace,identifier) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "endorsement_type" : "endorsement_type",
  "endrosement_value" : "endrosement_value"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Fetch Request History list for {namespace}/{identifier}
 * TBD
 *
 * namespace String ID Namespace
 * identifier String Identifier for the ID
 * count String Number of request to get (optional)
 * returns List
 **/
exports.get_request_history = function(namespace,identifier,count) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = [ {
  "request_body" : "request_body",
  "request_id" : "request_id"
}, {
  "request_body" : "request_body",
  "request_id" : "request_id"
} ];
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}

