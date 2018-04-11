'use strict';

var utils = require('../utils/writer.js');
var Default = require('../service/DefaultService');

module.exports.request_for_authentication = function request_for_authentication (req, res, next) {
  var namespace = req.swagger.params['namespace'].value;
  var identifier = req.swagger.params['identifier'].value;
  var requests = req.swagger.params['requests'].value;
  Default.request_for_authentication(namespace,identifier,requests)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
