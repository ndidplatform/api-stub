'use strict';

var utils = require('../utils/writer.js');
var Default = require('../service/DefaultService');

module.exports.get_request_data = function get_request_data (req, res, next) {
  var request_id = req.swagger.params['request_id'].value;
  Default.get_request_data(request_id)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.get_request_id_from_reference_number = function get_request_id_from_reference_number (req, res, next) {
  var reference_number = req.swagger.params['reference_number'].value;
  Default.get_request_id_from_reference_number(reference_number)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.get_request_status = function get_request_status (req, res, next) {
  var request_id = req.swagger.params['request_id'].value;
  Default.get_request_status(request_id)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.send_request_to_id = function send_request_to_id (req, res, next) {
  var namespace = req.swagger.params['namespace'].value;
  var identifier = req.swagger.params['identifier'].value;
  var request = req.swagger.params['request'].value;
  var timeout = req.swagger.params['timeout'].value;
  Default.send_request_to_id(namespace,identifier,request,timeout)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
