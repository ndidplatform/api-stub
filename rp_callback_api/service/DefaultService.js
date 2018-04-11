'use strict';


/**
 * Update from NDID
 * Update from NDID to RP, there's been an update to the status
 *
 * request_id String ID Namespace
 * status String Request status
 * no response value expected for this operation
 **/
exports.request_for_authentication = function(request_id,status) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}

