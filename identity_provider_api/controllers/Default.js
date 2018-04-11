'use strict';

var utils = require('../utils/writer.js');
var Default = require('../service/DefaultService');

module.exports.get_callback_url = function get_callback_url (req, res, next) {
  Default.get_callback_url()
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.respond_to_request = function respond_to_request (req, res, next) {
  var idp_response = req.swagger.params['idp_response'].value;
  Default.respond_to_request(idp_response)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.set_callback_url = function set_callback_url (req, res, next) {
  var url = req.swagger.params['url'].value;
  Default.set_callback_url(url)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
