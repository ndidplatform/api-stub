'use strict';

var utils = require('../utils/writer.js');
var Default = require('../service/DefaultService');

module.exports.add_accessor_method = function add_accessor_method (req, res, next) {
  var namespace = req.swagger.params['namespace'].value;
  var identifier = req.swagger.params['identifier'].value;
  var accessors = req.swagger.params['accessors'].value;
  Default.add_accessor_method(namespace,identifier,accessors)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.add_identifier = function add_identifier (req, res, next) {
  var namespace = req.swagger.params['namespace'].value;
  var identifier = req.swagger.params['identifier'].value;
  var new_identifiers = req.swagger.params['new_identifiers'].value;
  Default.add_identifier(namespace,identifier,new_identifiers)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.check_identity = function check_identity (req, res, next) {
  var namespace = req.swagger.params['namespace'].value;
  var identifier = req.swagger.params['identifier'].value;
  Default.check_identity(namespace,identifier)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.create_identity = function create_identity (req, res, next) {
  var identity = req.swagger.params['identity'].value;
  Default.create_identity(identity)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.endorse_identity = function endorse_identity (req, res, next) {
  var namespace = req.swagger.params['namespace'].value;
  var identifier = req.swagger.params['identifier'].value;
  var endorsements = req.swagger.params['endorsements'].value;
  Default.endorse_identity(namespace,identifier,endorsements)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.get_endorsement = function get_endorsement (req, res, next) {
  var namespace = req.swagger.params['namespace'].value;
  var identifier = req.swagger.params['identifier'].value;
  Default.get_endorsement(namespace,identifier)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.get_request_history = function get_request_history (req, res, next) {
  var namespace = req.swagger.params['namespace'].value;
  var identifier = req.swagger.params['identifier'].value;
  var count = req.swagger.params['count'].value;
  Default.get_request_history(namespace,identifier,count)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
