'use strict';


/**
 * Retrieve current callback url
 * asdf
 *
 * returns inline_response_201
 **/
exports.get_callback_url = function() {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
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
 * Respond to a request
 * asdf
 *
 * idp_response IDPResponse IDP_response
 * no response value expected for this operation
 **/
exports.respond_to_request = function(idp_response) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}


/**
 * Set callback url
 * asdf
 *
 * url String callback URL endpoint
 * no response value expected for this operation
 **/
exports.set_callback_url = function(url) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}

