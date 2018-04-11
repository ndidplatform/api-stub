'use strict';

var utils = require('../utils/writer.js');
var Default = require('../service/DefaultService');

module.exports.dpkiMasterSignPOST = function dpkiMasterSignPOST (req, res, next) {
  var sign_request = req.swagger.params['sign_request'].value;
  Default.dpkiMasterSignPOST(sign_request)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.dpkiSignPOST = function dpkiSignPOST (req, res, next) {
  var sign_request = req.swagger.params['sign_request'].value;
  Default.dpkiSignPOST(sign_request)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
