'use strict';

var utils = require('../utils/writer.js');
var Default = require('../service/DefaultService');

module.exports.serviceNamespaceIdentifierGET = function serviceNamespaceIdentifierGET (req, res, next) {
  var namespace = req.swagger.params['namespace'].value;
  var identifier = req.swagger.params['identifier'].value;
  var request = req.swagger.params['request'].value;
  Default.serviceNamespaceIdentifierGET(namespace,identifier,request)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
