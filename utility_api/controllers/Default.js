'use strict';

var utils = require('../utils/writer.js');
var Default = require('../service/DefaultService');

module.exports.get_as_id_list = function get_as_id_list (req, res, next) {
  var service_id = req.swagger.params['service_id'].value;
  Default.get_as_id_list(service_id)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.get_idp_list = function get_idp_list (req, res, next) {
  var min_ial = req.swagger.params['min_ial'].value;
  var min_aal = req.swagger.params['min_aal'].value;
  Default.get_idp_list(min_ial,min_aal)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.get_idp_list_for_identifier = function get_idp_list_for_identifier (req, res, next) {
  var namespace = req.swagger.params['namespace'].value;
  var identifier = req.swagger.params['identifier'].value;
  var min_ial = req.swagger.params['min_ial'].value;
  var min_aal = req.swagger.params['min_aal'].value;
  Default.get_idp_list_for_identifier(namespace,identifier,min_ial,min_aal)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
