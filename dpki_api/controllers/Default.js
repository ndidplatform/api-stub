'use strict';

var utils = require('../utils/writer.js');
var Default = require('../service/DefaultService');

module.exports.dpkiNodeCreatePOST = function dpkiNodeCreatePOST (req, res, next) {
  var node_data = req.swagger.params['node_data'].value;
  Default.dpkiNodeCreatePOST(node_data)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.dpkiNodeRegister_callbackPOST = function dpkiNodeRegister_callbackPOST (req, res, next) {
  var url = req.swagger.params['url'].value;
  Default.dpkiNodeRegister_callbackPOST(url)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.dpkiNodeRegister_callback_masterPOST = function dpkiNodeRegister_callback_masterPOST (req, res, next) {
  var url = req.swagger.params['url'].value;
  Default.dpkiNodeRegister_callback_masterPOST(url)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.dpkiNodeUpdatePOST = function dpkiNodeUpdatePOST (req, res, next) {
  var node_data = req.swagger.params['node_data'].value;
  Default.dpkiNodeUpdatePOST(node_data)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
