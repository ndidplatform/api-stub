'use strict';

var utils = require('../utils/writer.js');
var Default = require('../service/DefaultService');

module.exports.request_for_authentication = function request_for_authentication (req, res, next) {
  var request_id = req.swagger.params['request_id'].value;
  var status = req.swagger.params['status'].value;
  Default.request_for_authentication(request_id,status)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
