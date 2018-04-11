'use strict';


/**
 * Service X being offered by an AS to NDID
 * asdf
 *
 * namespace String 
 * identifier String 
 * request DataRequest data request format
 * returns DataMessage
 **/
exports.serviceNamespaceIdentifierGET = function(namespace,identifier,request) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "data" : "data"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}

