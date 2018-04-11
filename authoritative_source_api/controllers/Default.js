'use strict';

var utils = require('../utils/writer.js');
var Default = require('../service/DefaultService');

module.exports.get_service = function get_service (req, res, next) {
  var service_id = req.swagger.params['service_id'].value;
  Default.get_service(service_id)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.register_service = function register_service (req, res, next) {
  var service_id = req.swagger.params['service_id'].value;
  var service = req.swagger.params['service'].value;
  Default.register_service(service_id,service)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
