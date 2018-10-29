/**
 * CyberSource Flex API
 * Simple PAN tokenization service
 *
 * OpenAPI spec version: 0.0.1
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.2.3
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/V2paymentsOrderInformationAmountDetails', 'model/V2paymentsOrderInformationBillTo', 'model/V2paymentsOrderInformationInvoiceDetails', 'model/V2paymentsOrderInformationLineItems', 'model/V2paymentsOrderInformationShipTo', 'model/V2paymentsOrderInformationShippingDetails'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./V2paymentsOrderInformationAmountDetails'), require('./V2paymentsOrderInformationBillTo'), require('./V2paymentsOrderInformationInvoiceDetails'), require('./V2paymentsOrderInformationLineItems'), require('./V2paymentsOrderInformationShipTo'), require('./V2paymentsOrderInformationShippingDetails'));
  } else {
    // Browser globals (root is window)
    if (!root.CyberSource) {
      root.CyberSource = {};
    }
    root.CyberSource.V2paymentsOrderInformation = factory(root.CyberSource.ApiClient, root.CyberSource.V2paymentsOrderInformationAmountDetails, root.CyberSource.V2paymentsOrderInformationBillTo, root.CyberSource.V2paymentsOrderInformationInvoiceDetails, root.CyberSource.V2paymentsOrderInformationLineItems, root.CyberSource.V2paymentsOrderInformationShipTo, root.CyberSource.V2paymentsOrderInformationShippingDetails);
  }
}(this, function(ApiClient, V2paymentsOrderInformationAmountDetails, V2paymentsOrderInformationBillTo, V2paymentsOrderInformationInvoiceDetails, V2paymentsOrderInformationLineItems, V2paymentsOrderInformationShipTo, V2paymentsOrderInformationShippingDetails) {
  'use strict';




  /**
   * The V2paymentsOrderInformation model module.
   * @module model/V2paymentsOrderInformation
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>V2paymentsOrderInformation</code>.
   * @alias module:model/V2paymentsOrderInformation
   * @class
   */
  var exports = function() {
    var _this = this;







  };

  /**
   * Constructs a <code>V2paymentsOrderInformation</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/V2paymentsOrderInformation} obj Optional instance to populate.
   * @return {module:model/V2paymentsOrderInformation} The populated <code>V2paymentsOrderInformation</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('amountDetails')) {
        obj['amountDetails'] = V2paymentsOrderInformationAmountDetails.constructFromObject(data['amountDetails']);
      }
      if (data.hasOwnProperty('billTo')) {
        obj['billTo'] = V2paymentsOrderInformationBillTo.constructFromObject(data['billTo']);
      }
      if (data.hasOwnProperty('shipTo')) {
        obj['shipTo'] = V2paymentsOrderInformationShipTo.constructFromObject(data['shipTo']);
      }
      if (data.hasOwnProperty('lineItems')) {
        obj['lineItems'] = ApiClient.convertToType(data['lineItems'], [V2paymentsOrderInformationLineItems]);
      }
      if (data.hasOwnProperty('invoiceDetails')) {
        obj['invoiceDetails'] = V2paymentsOrderInformationInvoiceDetails.constructFromObject(data['invoiceDetails']);
      }
      if (data.hasOwnProperty('shippingDetails')) {
        obj['shippingDetails'] = V2paymentsOrderInformationShippingDetails.constructFromObject(data['shippingDetails']);
      }
    }
    return obj;
  }

  /**
   * @member {module:model/V2paymentsOrderInformationAmountDetails} amountDetails
   */
  exports.prototype['amountDetails'] = undefined;
  /**
   * @member {module:model/V2paymentsOrderInformationBillTo} billTo
   */
  exports.prototype['billTo'] = undefined;
  /**
   * @member {module:model/V2paymentsOrderInformationShipTo} shipTo
   */
  exports.prototype['shipTo'] = undefined;
  /**
   * @member {Array.<module:model/V2paymentsOrderInformationLineItems>} lineItems
   */
  exports.prototype['lineItems'] = undefined;
  /**
   * @member {module:model/V2paymentsOrderInformationInvoiceDetails} invoiceDetails
   */
  exports.prototype['invoiceDetails'] = undefined;
  /**
   * @member {module:model/V2paymentsOrderInformationShippingDetails} shippingDetails
   */
  exports.prototype['shippingDetails'] = undefined;



  return exports;
}));

