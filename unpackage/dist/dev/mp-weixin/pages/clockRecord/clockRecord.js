(global["webpackJsonp"] = global["webpackJsonp"] || []).push([["pages/clockRecord/clockRecord"],{

/***/ 133:
/*!*******************************************************************************************************************************!*\
  !*** G:/vue-init/uniapp-template-init/Infectious-disease-identification/main.js?{"page":"pages%2FclockRecord%2FclockRecord"} ***!
  \*******************************************************************************************************************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";
/* WEBPACK VAR INJECTION */(function(wx, createPage) {

var _interopRequireDefault = __webpack_require__(/*! @babel/runtime/helpers/interopRequireDefault */ 4);
__webpack_require__(/*! uni-pages */ 26);
var _vue = _interopRequireDefault(__webpack_require__(/*! vue */ 25));
var _clockRecord = _interopRequireDefault(__webpack_require__(/*! ./pages/clockRecord/clockRecord.vue */ 134));
// @ts-ignore
wx.__webpack_require_UNI_MP_PLUGIN__ = __webpack_require__;
createPage(_clockRecord.default);
/* WEBPACK VAR INJECTION */}.call(this, __webpack_require__(/*! ./node_modules/@dcloudio/uni-mp-weixin/dist/wx.js */ 1)["default"], __webpack_require__(/*! ./node_modules/@dcloudio/uni-mp-weixin/dist/index.js */ 2)["createPage"]))

/***/ }),

/***/ 134:
/*!************************************************************************************************************!*\
  !*** G:/vue-init/uniapp-template-init/Infectious-disease-identification/pages/clockRecord/clockRecord.vue ***!
  \************************************************************************************************************/
/*! no static exports found */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _clockRecord_vue_vue_type_template_id_33f46406___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./clockRecord.vue?vue&type=template&id=33f46406& */ 135);
/* harmony import */ var _clockRecord_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./clockRecord.vue?vue&type=script&lang=js& */ 137);
/* harmony reexport (unknown) */ for(var __WEBPACK_IMPORT_KEY__ in _clockRecord_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__) if(["default"].indexOf(__WEBPACK_IMPORT_KEY__) < 0) (function(key) { __webpack_require__.d(__webpack_exports__, key, function() { return _clockRecord_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__[key]; }) }(__WEBPACK_IMPORT_KEY__));
/* harmony import */ var _clockRecord_vue_vue_type_style_index_0_lang_scss___WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./clockRecord.vue?vue&type=style&index=0&lang=scss& */ 139);
/* harmony import */ var _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib/runtime/componentNormalizer.js */ 32);

var renderjs





/* normalize component */

var component = Object(_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _clockRecord_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__["default"],
  _clockRecord_vue_vue_type_template_id_33f46406___WEBPACK_IMPORTED_MODULE_0__["render"],
  _clockRecord_vue_vue_type_template_id_33f46406___WEBPACK_IMPORTED_MODULE_0__["staticRenderFns"],
  false,
  null,
  null,
  null,
  false,
  _clockRecord_vue_vue_type_template_id_33f46406___WEBPACK_IMPORTED_MODULE_0__["components"],
  renderjs
)

component.options.__file = "pages/clockRecord/clockRecord.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ 135:
/*!*******************************************************************************************************************************************!*\
  !*** G:/vue-init/uniapp-template-init/Infectious-disease-identification/pages/clockRecord/clockRecord.vue?vue&type=template&id=33f46406& ***!
  \*******************************************************************************************************************************************/
/*! exports provided: render, staticRenderFns, recyclableRender, components */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_17_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_template_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_uni_app_loader_page_meta_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_clockRecord_vue_vue_type_template_id_33f46406___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-preprocess-loader??ref--17-0!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/template.js!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-uni-app-loader/page-meta.js!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib??vue-loader-options!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/style.js!./clockRecord.vue?vue&type=template&id=33f46406& */ 136);
/* harmony reexport (safe) */ __webpack_require__.d(__webpack_exports__, "render", function() { return _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_17_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_template_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_uni_app_loader_page_meta_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_clockRecord_vue_vue_type_template_id_33f46406___WEBPACK_IMPORTED_MODULE_0__["render"]; });

/* harmony reexport (safe) */ __webpack_require__.d(__webpack_exports__, "staticRenderFns", function() { return _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_17_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_template_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_uni_app_loader_page_meta_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_clockRecord_vue_vue_type_template_id_33f46406___WEBPACK_IMPORTED_MODULE_0__["staticRenderFns"]; });

/* harmony reexport (safe) */ __webpack_require__.d(__webpack_exports__, "recyclableRender", function() { return _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_17_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_template_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_uni_app_loader_page_meta_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_clockRecord_vue_vue_type_template_id_33f46406___WEBPACK_IMPORTED_MODULE_0__["recyclableRender"]; });

/* harmony reexport (safe) */ __webpack_require__.d(__webpack_exports__, "components", function() { return _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_17_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_template_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_uni_app_loader_page_meta_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_clockRecord_vue_vue_type_template_id_33f46406___WEBPACK_IMPORTED_MODULE_0__["components"]; });



/***/ }),

/***/ 136:
/*!*******************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-preprocess-loader??ref--17-0!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/template.js!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-uni-app-loader/page-meta.js!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib??vue-loader-options!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/style.js!G:/vue-init/uniapp-template-init/Infectious-disease-identification/pages/clockRecord/clockRecord.vue?vue&type=template&id=33f46406& ***!
  \*******************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/*! exports provided: render, staticRenderFns, recyclableRender, components */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "render", function() { return render; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "staticRenderFns", function() { return staticRenderFns; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "recyclableRender", function() { return recyclableRender; });
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "components", function() { return components; });
var components
try {
  components = {
    uniCalendar: function () {
      return Promise.all(/*! import() | uni_modules/uni-calendar/components/uni-calendar/uni-calendar */[__webpack_require__.e("common/vendor"), __webpack_require__.e("uni_modules/uni-calendar/components/uni-calendar/uni-calendar")]).then(__webpack_require__.bind(null, /*! @/uni_modules/uni-calendar/components/uni-calendar/uni-calendar.vue */ 242))
    },
    uniCustomTabbar: function () {
      return __webpack_require__.e(/*! import() | components/uni-custom-tabbar/uni-custom-tabbar */ "components/uni-custom-tabbar/uni-custom-tabbar").then(__webpack_require__.bind(null, /*! @/components/uni-custom-tabbar/uni-custom-tabbar.vue */ 217))
    },
  }
} catch (e) {
  if (
    e.message.indexOf("Cannot find module") !== -1 &&
    e.message.indexOf(".vue") !== -1
  ) {
    console.error(e.message)
    console.error("1. 排查组件名称拼写是否正确")
    console.error(
      "2. 排查组件是否符合 easycom 规范，文档：https://uniapp.dcloud.net.cn/collocation/pages?id=easycom"
    )
    console.error(
      "3. 若组件不符合 easycom 规范，需手动引入，并在 components 中注册该组件"
    )
  } else {
    throw e
  }
}
var render = function () {
  var _vm = this
  var _h = _vm.$createElement
  var _c = _vm._self._c || _h
  var g0 = _vm.diseases.length
  var g1 = _vm.diseases.length
  _vm.$mp.data = Object.assign(
    {},
    {
      $root: {
        g0: g0,
        g1: g1,
      },
    }
  )
}
var recyclableRender = false
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ 137:
/*!*************************************************************************************************************************************!*\
  !*** G:/vue-init/uniapp-template-init/Infectious-disease-identification/pages/clockRecord/clockRecord.vue?vue&type=script&lang=js& ***!
  \*************************************************************************************************************************************/
/*! no static exports found */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_babel_loader_lib_index_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_13_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_script_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_clockRecord_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!./node_modules/babel-loader/lib!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-preprocess-loader??ref--13-1!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/script.js!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib??vue-loader-options!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/style.js!./clockRecord.vue?vue&type=script&lang=js& */ 138);
/* harmony import */ var _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_babel_loader_lib_index_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_13_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_script_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_clockRecord_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_babel_loader_lib_index_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_13_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_script_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_clockRecord_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__);
/* harmony reexport (unknown) */ for(var __WEBPACK_IMPORT_KEY__ in _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_babel_loader_lib_index_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_13_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_script_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_clockRecord_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__) if(["default"].indexOf(__WEBPACK_IMPORT_KEY__) < 0) (function(key) { __webpack_require__.d(__webpack_exports__, key, function() { return _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_babel_loader_lib_index_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_13_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_script_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_clockRecord_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__[key]; }) }(__WEBPACK_IMPORT_KEY__));
 /* harmony default export */ __webpack_exports__["default"] = (_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_babel_loader_lib_index_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_13_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_script_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_clockRecord_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0___default.a); 

/***/ }),

/***/ 138:
/*!********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-preprocess-loader??ref--13-1!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/script.js!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib??vue-loader-options!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/style.js!G:/vue-init/uniapp-template-init/Infectious-disease-identification/pages/clockRecord/clockRecord.vue?vue&type=script&lang=js& ***!
  \********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";
/* WEBPACK VAR INJECTION */(function(wx, uni) {

var _interopRequireDefault = __webpack_require__(/*! @babel/runtime/helpers/interopRequireDefault */ 4);
Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = void 0;
var _regenerator = _interopRequireDefault(__webpack_require__(/*! @babel/runtime/regenerator */ 42));
var _asyncToGenerator2 = _interopRequireDefault(__webpack_require__(/*! @babel/runtime/helpers/asyncToGenerator */ 44));
var _toConsumableArray2 = _interopRequireDefault(__webpack_require__(/*! @babel/runtime/helpers/toConsumableArray */ 18));
var _defineProperty2 = _interopRequireDefault(__webpack_require__(/*! @babel/runtime/helpers/defineProperty */ 11));
var _typeof2 = _interopRequireDefault(__webpack_require__(/*! @babel/runtime/helpers/typeof */ 13));
var _clock = __webpack_require__(/*! ../../services/clock */ 90);
var _index = _interopRequireDefault(__webpack_require__(/*! ../../store/index.ts */ 63));
var _methods;
/**
 * 获取任意时间
 */
function getDate(date) {
  var AddDayCount = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 0;
  if (!date) {
    date = new Date();
  }
  if ((0, _typeof2.default)(date) !== 'object') {
    date = date.replace(/-/g, '/');
  }
  var dd = new Date(date);
  dd.setDate(dd.getDate() + AddDayCount); // 获取AddDayCount天后的日期

  var y = dd.getFullYear();
  var m = dd.getMonth() + 1 < 10 ? '0' + (dd.getMonth() + 1) : dd.getMonth() + 1; // 获取当前月份的日期，不足10补0
  var d = dd.getDate() < 10 ? '0' + dd.getDate() : dd.getDate(); // 获取当前几号，不足10补0
  return {
    fullDate: y + '-' + m + '-' + d,
    year: y,
    month: m,
    date: d,
    day: dd.getDay()
  };
}
var _default = {
  data: function data() {
    var _ref;
    return _ref = {
      dateForm: {
        date: '',
        diagnose: '',
        examine: '',
        health: ''
      },
      dateForms: {
        date: ''
      },
      reportForm: {
        isSerologicalTestDone: false,
        isVirusAntigenTestDone: false,
        isVirusNucleicAcidTestDone: false,
        isVirusCultureIsolationDone: false,
        labTestFiles: [] // 新增的 labTestFileIds 数组
      },

      todayInfo: '',
      yearMonth: '',
      checkFormData: {
        stoolReport: [],
        // 粪便检查报告文件数组
        bloodReport: [],
        // 血液检查报告文件数组
        imageReport: [],
        ultrasoundReport: [],
        urineReport: [],
        medicalRecord: []
      },
      formData: {
        date: '',
        stoolReport: [],
        // 粪便检查报告文件数组
        vomitReport: [
          // 可以添加更多文件
        ],
        // 呕吐物检查报告文件数组
        bloodReport: [
          // 可以添加更多文件
        ] // 血液检查报告文件数组
      },

      showCalendar: false,
      info: {
        lunar: true,
        range: true,
        insert: false,
        selected: [] // 日期打卡记录
      },

      tempRecords: [],
      records: {},
      diseases: [{
        name: '新型冠状病毒',
        rate: '90%',
        link: '/pages/healthProtection/healthProtection'
      }, {
        name: '流感',
        rate: '80%',
        link: '/pages/healthProtection/healthProtection'
      }, {
        name: '鼠疫',
        rate: '75%',
        link: '/pages/healthProtection/healthProtection'
      }, {
        name: '感染性腹泻',
        rate: '60%',
        link: '/pages/healthProtection/healthProtection'
      }, {
        name: '炭疽',
        rate: '50%',
        link: '/pages/healthProtection/healthProtection'
      }],
      calendarPages: true,
      currentDateNow: '',
      examineReport1: false,
      examineReport2: false,
      sonpages: 1,
      pages: 1
    }, (0, _defineProperty2.default)(_ref, "formData", {
      date: '',
      stoolReport: [
        // 可以添加更多文件
      ],
      // 粪便检查报告文件数组
      vomitReport: [
        // 可以添加更多文件
      ],
      // 呕吐物检查报告文件数组
      bloodReport: [
        // 可以添加更多文件
      ] // 血液检查报告文件数组
    }), (0, _defineProperty2.default)(_ref, "AllCheckInfo", {
      diagnosisResultsEmpDTO: {
        admissionDate: '',
        admissionSymptomsAndSigns: '',
        anthraxSubtype: '',
        deathDate: null,
        dischargeDate: '',
        discoveryMethod: '',
        diseaseOutcome: '',
        diseaseType: '',
        hospitalName: '',
        otherDiscoveryMethodName: null,
        otherDiseaseName: null,
        otherRegistrationDetails: null,
        plagueSubtype: null,
        registrationClassification: ''
      },
      diagnosiscirculatorysymptoms: {
        hasArrhythmia: false,
        hasChestPain: false,
        hasPalpitation: false,
        hasRapidPulse: false
      },
      diagnosiscomplications: {
        hasAbdominalHemorrhage: false,
        hasAcuteKidneyInjury: false,
        hasAcuteLungInjury: false,
        hasAnemia: false,
        hasArds: false,
        hasBacteremia: false,
        hasBacterialPneumonia: false,
        hasBrainEdema: false,
        hasBronchitis: false,
        hasBronchopneumonia: false,
        hasCoagulationDisorder: false,
        hasComplications: false,
        hasDic: false,
        hasElectrolyteImbalance: false,
        hasEncephalitis: false,
        hasEpilepsy: false,
        hasFluidOverload: false,
        hasFungalPneumonia: false,
        hasGastrointestinalBleeding: false,
        hasHematemesis: false,
        hasHemolysis: false,
        hasHemoptysis: false,
        hasHepatosplenomegaly: false,
        hasHus: false,
        hasHypertensiveEncephalopathy: false,
        hasHypoxemia: false,
        hasIntracranialHemorrhage: false,
        hasLiverDamage: false,
        hasMelena: false,
        hasMeningitis: false,
        hasMetabolicAcidosis: false,
        hasMultipleOrganFailure: false,
        hasMyocardialDamage: false,
        hasMyocarditis: false,
        hasNasalBleeding: false,
        hasPulmonaryEdema: false,
        hasRespiratoryFailure: false,
        hasSecondaryInfection: false,
        hasSepticShock: false,
        hasSpontaneousRenalRupture: false,
        hasSuperinfection: false,
        hasToxicHepatitis: false,
        hasToxicShock: false,
        hasVaginalBleeding: false,
        hasViralPneumonia: false,
        otherComplications: false,
        otherComplicationsName: null
      },
      diagnosisdigestivesymptoms: {
        diarrheaFrequencyGEThreeTimesPerDay: false,
        diarrheaMode: null,
        diarrheaTimesPerDay: 0,
        hasAbdominalDistension: false,
        hasAbdominalPain: false,
        hasAppetiteLoss: false,
        hasBorborygmus: false,
        hasConstipation: false,
        hasNausea: false,
        hasOliguriaOrAnuria: false,
        hasStoolType1: false,
        hasStoolType2: false,
        hasUpperAbdominalDiscomfort: false,
        hasVomiting: false,
        stoolAmount: '',
        stoolOdor: '',
        stoolType1Detail: null,
        stoolType2Detail: null,
        vomitingMode: null
      },
      diagnosisexaminations: {
        hasBloodTest: false,
        hasImaging: false,
        hasStoolTest: false,
        hasUltrasound: false,
        hasUrineTest: false
      },
      diagnosisgeneralsymptoms: {
        bleedingCavity: '',
        bleedingVolume: '',
        dehydrationSeverity: '',
        feverDuration: '',
        hasBleeding: false,
        hasBloodBlisters: false,
        hasBloodPressureDrop: false,
        hasBrudzinskiSign: false,
        hasChills: false,
        hasCongestiveOrPetechialRash: false,
        hasCyanosis: false,
        hasDehydration: false,
        hasEdema: false,
        hasExhaustion: false,
        hasFatigue: false,
        hasFever: false,
        hasHeadache: false,
        hasHighIntracranialPressure: false,
        hasItchyRash: false,
        hasJointPain: false,
        hasKernigSign: false,
        hasLymphNodeSwelling: false,
        hasLymphNodeSwellingArmpit: false,
        hasLymphNodeSwellingElbow: false,
        hasLymphNodeSwellingGroin: false,
        hasLymphNodeSwellingNeck: false,
        hasLymphNodeSwellingPopliteal: false,
        hasLymphNodeSwellingSubclavian: false,
        hasMusclePain: false,
        hasNightSweats: false,
        hasOtherOrganDamage: false,
        hasPainfulRedRash: false,
        hasPlasmaLeakage: false,
        hasSevereAnemia: false,
        hasShockSyndrome: false,
        hasSkinUlcer: false,
        hasSubcutaneousAndMucosalBleedingCavity: false,
        hasSubcutaneousAndMucosalBleedingSpots: false,
        hasSweating: false,
        hasTurbidCerebrospinalFluid: false,
        hasWeightLoss: false,
        hasWhiteBloodCells: false,
        highestTemperature: '',
        itchyRashOnBackOfHands: false,
        itchyRashOnFace: false,
        itchyRashOnFeet: false,
        itchyRashOnFingers: false,
        itchyRashOnLowerLimbs: false,
        itchyRashOnOther: false,
        itchyRashOnUpperLimbs: false,
        lowestBloodPressure: ''
      },
      diagnosislocalsymptoms: {
        hasConjunctivalCongestion: false,
        hasConjunctivitis: false,
        hasEyelidEdema: false,
        hasFacialRednessOrPallor: false,
        hasGumBleeding: false,
        hasHepatomegaly: false,
        hasKidneyPain: false,
        hasLipCyanosis: false,
        hasLowBackPain: false,
        hasMucosalBleedingInOralOrNasal: false,
        hasOrbitalPain: false,
        hasRednessNeckChest: false,
        hasSkinBleedingPoints: false,
        hasSkinPallorOrBruising: false,
        hasSmellTasteLoss: false,
        hasSplenomegaly: false
      },
      diagnosisneurologicalsymptoms: {
        hasComa: false,
        hasConsciousnessDisorder: false,
        hasConvulsion: false,
        hasDelirium: false,
        hasDizziness: false,
        hasDrowsiness: false,
        hasHearingLoss: false,
        hasMuscleParalysisLimbs: false,
        hasMuscleParalysisNeckShoulder: false,
        hasRestlessness: false,
        hasSpeechDisorder: false,
        hasStiffNeck: false,
        hasSwallowingDifficulty: false,
        hasTinnitus: false
      },
      diagnosisothersymptoms: {
        hasCalfMusclePain: false,
        hasForcedPosture: false,
        hasOtherSymptoms: false,
        hasPeriodicAttack: false,
        hasRapidProgress: false,
        hasSuddenOnset: false,
        otherSymptomsName: null
      },
      diagnosispersonalinfo: {
        bloodOxygen: '',
        bloodPressure: '',
        bloodSugarFasting: '',
        bloodSugarRandom: '',
        cholesterolLipid: '',
        cholesterolTotal: '',
        hdl: '',
        height: '',
        ldl: '',
        triglycerides: '',
        weight: ''
      },
      diagnosisrespiratorysymptoms: {
        hasBloodySputum: false,
        hasBreathingDifficulty: false,
        hasBronchitis: false,
        hasChestTightness: false,
        hasCough: false,
        hasDryThroat: false,
        hasHemoptysis: false,
        hasHoarseness: false,
        hasLungRales: false,
        hasNasalCongestion: false,
        hasPharyngealHyperemia: false,
        hasRunnyNose: false,
        hasShortnessOfBreath: false,
        hasSoreThroat: false,
        hasSputum: false
      },
      fileIdList: []
    }), (0, _defineProperty2.default)(_ref, "temporaryChoice1", false), (0, _defineProperty2.default)(_ref, "temporaryChoice2", false), (0, _defineProperty2.default)(_ref, "otherChoices", false), (0, _defineProperty2.default)(_ref, "SeafoodChoices", false), _ref;
  },
  onReady: function onReady() {
    var _this = this;
    this.$nextTick(function () {
      _this.showCalendar = true;
    });

    // 模拟请求异步同步数据
    setTimeout(function () {
      _this.info.date = getDate(new Date(), -30).fullDate;
      _this.info.startDate = getDate(new Date(), -60).fullDate;
      _this.info.endDate = getDate(new Date(), 30).fullDate;
      _this.info.selected = Object.keys(_this.records).map(function (date) {
        return {
          date: date,
          info: _this.records[date]
        };
      });
    }, 300);
  },
  mounted: function mounted() {
    var _this2 = this;
    this.initMonthSwitch();
    (0, _clock.clockRecordService)(this.yearMonth).then(function (res) {
      console.log(res.data);
      _this2.tempRecords = res.data; // 保存所有的记录数据
      res.data.forEach(function (item) {
        _this2.updateRecords(item);
      });
    }).catch(function (error) {
      return _this2.submitError(error);
    });
    var diseasescoring = _index.default.state.diseasescoring;
    var formattedDiseases = diseasescoring.map(function (item) {
      return {
        diseaseTypeName: item.diseaseTypeName,
        score: item.score,
        link: '/pages/healthProtection/healthProtection'
      };
    });
    // 对 formattedDiseases 进行排序
    formattedDiseases.sort(function (a, b) {
      return b.score - a.score;
    });
    // 将排序后的数据赋值给 diseases
    this.diseases = formattedDiseases;
  },
  methods: (_methods = {
    preventDelete: function preventDelete() {
      // 防止删除操作
      console.log("删除操作被禁止");
      return false; // 返回 false 以防止删除
    },
    formatDate: function formatDate(date) {
      var year = date.getFullYear();
      var month = String(date.getMonth() + 1).padStart(2, '0'); // 月份从0开始，所以加1
      var day = String(date.getDate()).padStart(2, '0');
      return "".concat(year, "-").concat(month, "-").concat(day);
    },
    checkSuccess: function checkSuccess(response) {
      console.log("666666666", response);
      // 接收到的数据进行排序
      var sortedDiseases = response.data.diseasescoringList.sort(function (a, b) {
        return b.score - a.score;
      });
      this.diseases = sortedDiseases;
    },
    checkError: function checkError(response) {
      console.log(response);
    },
    updateRecords: function updateRecords(data) {
      var _this3 = this;
      var date = data.date,
        health = data.health,
        diagnose = data.diagnose,
        examine = data.examine;
      var key = date;
      if (!this.records[key]) {
        this.records[key] = [];
      }
      this.records[key][0] = health ? 1 : 0;
      this.records[key][1] = diagnose ? 1 : 0;
      this.records[key][2] = examine ? 1 : 0;
      console.log('Updated records:', this.records);
      this.$nextTick(function () {
        _this3.showCalendar = true;
      });

      // 模拟请求异步同步数据
      setTimeout(function () {
        _this3.info.date = getDate(new Date(), -30).fullDate;
        _this3.info.startDate = getDate(new Date(), -60).fullDate;
        _this3.info.endDate = getDate(new Date(), 30).fullDate;
        _this3.info.selected = Object.keys(_this3.records).map(function (date) {
          return {
            date: date,
            info: _this3.records[date]
          };
        });
      }, 300);
    },
    initMonthSwitch: function initMonthSwitch() {
      // 获取当前日期
      var currentDate = new Date();
      var currentYear = currentDate.getFullYear();
      var currentMonth = String(currentDate.getMonth() + 1).padStart(2, '0'); // 确保月份为两位数
      var currentDay = String(currentDate.getDate()).padStart(2, '0'); // 确保日期为两位数

      // 存储当前的年月信息
      this.currentYear = currentYear;
      this.currentMonth = currentMonth;
      this.currentDay = currentDay;
      this.yearMonth = "".concat(currentYear, "-").concat(currentMonth); // 拼接成 YYYY-MM 格式
      this.todayInfo = "".concat(currentYear, "-").concat(currentMonth, "-").concat(currentDay); // 拼接成 YYYY-MM-DD 格式

      console.log(this.todayInfo);
      console.log('nownownownownow', this.yearMonth);
    },
    isImage: function isImage(filePath) {
      var extension = filePath.split('.').pop().toLowerCase();
      return ['jpg', 'png'].includes(extension);
    },
    onDateChange: function onDateChange(e) {
      this.formData.date = e.detail.value;
    },
    chooseFile: function chooseFile(reportType) {
      var _this4 = this;
      wx.chooseMessageFile({
        count: 3,
        // 最多可以选择3个文件
        type: 'file',
        extension: ['pdf', 'jpg', 'png'],
        // 支持的文件类型
        success: function success(res) {
          var files = res.tempFiles.map(function (file) {
            return {
              path: file.path,
              name: file.name
            };
          });
          _this4.formData[reportType] = [].concat((0, _toConsumableArray2.default)(_this4.formData[reportType]), (0, _toConsumableArray2.default)(files));
        }
      });
    },
    uploadFile: function uploadFile(filePath, reportType, index) {
      var _this5 = this;
      return new Promise(function (resolve, reject) {
        wx.uploadFile({
          url: 'https://ds.sccdc.cn/railwayemployee/labtest/file',
          // 替换为你的接口地址
          filePath: filePath,
          name: 'file',
          // 根据你的接口定义
          formData: {
            reportType: reportType,
            // 上传的报告类型
            index: index // 当前文件的索引
          },

          header: {
            'Content-Type': 'application/x-www-form-urlencoded'
          },
          success: function success(res) {
            console.log("\u6587\u4EF6\u4E0A\u4F20\u6210\u529F: ".concat(res));
            var fileId = parseInt(res.data); // 假设返回的是一个数字
            _this5.reportForm.labTestFiles.push(fileId);
            resolve(res);
          },
          fail: function fail(err) {
            console.error("\u6587\u4EF6\u4E0A\u4F20\u5931\u8D25: ".concat(err));
            reject(err);
          }
        });
      });
    },
    change: function change(e) {
      var _this6 = this;
      // 检查所选日期是否在未来
      var currentDate = getDate(new Date()).fullDate;
      var selectedDate = e.fulldate;
      console.log('change 返回:', e, "66666666", currentDate, selectedDate);
      this.currentDateNow = e.fulldate;
      if (selectedDate > currentDate) {
        this.showUnmarkedPopup(selectedDate);
      } else {
        this.calendarPages = !this.calendarPages;
      }
      this.todayInfo = selectedDate;

      // 查找与选中日期匹配的记录
      var selectedRecord = this.tempRecords.find(function (record) {
        return record.date === _this6.todayInfo;
      });
      if (selectedRecord) {
        // 如果找到匹配的记录，赋值 diagnose, examine, health 属性
        this.dateForm.diagnose = selectedRecord.diagnose;
        this.dateForm.examine = selectedRecord.examine;
        this.dateForm.health = selectedRecord.health;
      } else {
        // 如果没有找到匹配的记录，设置默认值
        this.dateForm.diagnose = false;
        this.dateForm.examine = false;
        this.dateForm.health = false;
      }
      console.log('测试', this.dateForm);
      this.dateForms.date = this.todayInfo;
      (0, _clock.clockSituationService)(this.dateForms).then(function (res) {
        return _this6.checkSuccess(res);
      }).catch(function (error) {
        return _this6.checkError(error);
      });
    },
    showUnmarkedPopup: function showUnmarkedPopup(date) {
      uni.showToast({
        title: "\u65E5\u671F ".concat(date, " \u5C1A\u672A\u6253\u5361"),
        icon: 'none',
        duration: 2000
      });
    },
    monthSwitch: function monthSwitch(e) {
      var _this7 = this;
      console.log('monthSwitch 返回:', e);
      this.yearMonth = e.year + '-' + e.month;
      (0, _clock.clockRecordService)(this.yearMonth).then(function (res) {
        console.log(res.data);
        res.data.forEach(function (item) {
          _this7.updateRecords(item);
        });
      }).catch(function (error) {
        return _this7.submitError(error);
      });
    },
    navigateToDay: function navigateToDay(day) {
      var selectedDate = day.fullDate;
      wx.navigateTo({
        url: "/pages/targetPage/targetPage?date=".concat(selectedDate)
      });
    }
  }, (0, _defineProperty2.default)(_methods, "onDateChange", function onDateChange(e) {
    this.formData.date = e.detail.value;
  }), (0, _defineProperty2.default)(_methods, "chooseFile", function chooseFile(reportType) {
    var _this8 = this;
    wx.chooseMessageFile({
      count: 3,
      // 最多可以选择3个文件
      type: 'file',
      extension: ['pdf', 'jpg', 'png'],
      // 支持的文件类型
      success: function success(res) {
        var files = res.tempFiles.map(function (file) {
          return {
            path: file.path,
            name: file.name
          };
        });
        _this8.formData[reportType] = [].concat((0, _toConsumableArray2.default)(_this8.formData[reportType]), (0, _toConsumableArray2.default)(files));
      }
    });
  }), (0, _defineProperty2.default)(_methods, "uploadFile", function uploadFile(filePath, reportType, index) {
    var _this9 = this;
    return new Promise(function (resolve, reject) {
      wx.uploadFile({
        url: 'https://ds.sccdc.cn/railwayemployee/diagnosis/file',
        // 替换为你的接口地址
        filePath: filePath,
        name: 'file',
        // 根据你的接口定义
        formData: {
          reportType: reportType,
          // 上传的报告类型
          index: index // 当前文件的索引
        },

        header: {
          'Content-Type': 'application/x-www-form-urlencoded'
        },
        success: function success(res) {
          console.log("\u6587\u4EF6\u4E0A\u4F20\u6210\u529F: ".concat(res));
          var fileId = parseInt(res.data); // 假设返回的是一个数字
          _this9.AllCheckInfo.fileIdList.push(fileId);
          resolve(res);
        },
        fail: function fail(err) {
          console.error("\u6587\u4EF6\u4E0A\u4F20\u5931\u8D25: ".concat(err));
          reject(err);
        }
      });
    });
  }), (0, _defineProperty2.default)(_methods, "handleSubmits", function handleSubmits() {
    var _this10 = this;
    wx.showModal({
      content: '提交后不可修改，是否确认提交？',
      cancelText: '取消',
      confirmText: '确认提交',
      success: function () {
        var _success = (0, _asyncToGenerator2.default)( /*#__PURE__*/_regenerator.default.mark(function _callee(res) {
          var uploadPromises;
          return _regenerator.default.wrap(function _callee$(_context) {
            while (1) {
              switch (_context.prev = _context.next) {
                case 0:
                  if (!res.confirm) {
                    _context.next = 19;
                    break;
                  }
                  // 上传所有报告的文件
                  uploadPromises = []; // 上传粪便检查报告
                  _this10.formData.stoolReport.forEach(function (file, index) {
                    uploadPromises.push(_this10.uploadFile(file.path, '粪便检查报告', index));
                  });

                  // 上传呕吐物检查报告
                  _this10.formData.vomitReport.forEach(function (file, index) {
                    uploadPromises.push(_this10.uploadFile(file.path, '呕吐物检查报告', index));
                  });

                  // 上传血液检查报告
                  _this10.formData.bloodReport.forEach(function (file, index) {
                    uploadPromises.push(_this10.uploadFile(file.path, '血液检查报告', index));
                  });
                  _context.prev = 5;
                  _context.next = 8;
                  return Promise.all(uploadPromises);
                case 8:
                  console.log('所有文件上传完成');
                  // 提交表单
                  diagnosisSubmitService(_this10.AllCheckInfo).then(function (res) {
                    return _this10.submitSuccess(res);
                  }).catch(function (error) {
                    return _this10.submitError(error);
                  });
                  wx.showToast({
                    title: '提交成功',
                    icon: 'success',
                    duration: 2000
                  });
                  _context.next = 17;
                  break;
                case 13:
                  _context.prev = 13;
                  _context.t0 = _context["catch"](5);
                  console.error('上传过程出现错误: ', _context.t0);
                  wx.showToast({
                    title: '提交失败',
                    icon: 'none',
                    duration: 2000
                  });
                case 17:
                  _context.next = 20;
                  break;
                case 19:
                  if (res.cancel) {
                    console.log('用户点击取消');
                  }
                case 20:
                case "end":
                  return _context.stop();
              }
            }
          }, _callee, null, [[5, 13]]);
        }));
        function success(_x) {
          return _success.apply(this, arguments);
        }
        return success;
      }()
    });
  }), (0, _defineProperty2.default)(_methods, "submitError", function submitError(error) {}), (0, _defineProperty2.default)(_methods, "submitSuccess", function submitSuccess(response) {
    console.log("注意注意", response);
    if (response.code == 1) {
      uni.showToast({
        title: '信息提交成功',
        icon: 'success',
        duration: 1800
      });
    } else {
      uni.showModal({
        title: '信息提交失败',
        content: response.msg,
        showCancel: false,
        confirmText: '确定',
        success: function success(res) {}
      });
    }
  }), (0, _defineProperty2.default)(_methods, "handleReport", function handleReport(propName) {
    console.log("Changing ".concat(propName, ":"), this.reportForm[propName]);
    this.reportForm[propName] = !this.reportForm[propName];
  }), (0, _defineProperty2.default)(_methods, "isImage", function isImage(filePath) {
    var extension = filePath.split('.').pop().toLowerCase();
    return ['jpg', 'png'].includes(extension);
  }), (0, _defineProperty2.default)(_methods, "handlereturn", function handlereturn() {
    if (this.pages > 1) {
      this.pages--;
    } else if (this.pages == 1 && this.sonpages > 1) {
      this.sonpages--;
    }
  }), (0, _defineProperty2.default)(_methods, "checkExamineData", function checkExamineData() {
    console.log('日期准备', this.todayInfo);
    // 检查 examine 的状态
    if (this.dateForm.diagnose === false) {
      // 如果 examine 为 false，显示提示信息并阻止跳转
      uni.showToast({
        title: '用户当日未提交检测数据',
        icon: 'none',
        duration: 2000
      });
    } else {
      // 如果 examine 为 true，跳转到 diagnosis 页面并携带日期参数
      var dateParam = this.todayInfo;
      uni.navigateTo({
        url: "/pages/diagnosis/diagnosis?date=".concat(dateParam)
      });
    }
  }), (0, _defineProperty2.default)(_methods, "checkHealthData", function checkHealthData() {
    console.log('日期准备', this.todayInfo);
    if (this.dateForm.health === true) {
      uni.showToast({
        title: '用户打卡健康，未提交当日症状',
        icon: 'none',
        duration: 2000
      });
    } else {
      // 如果 examine 为 true，跳转到 diagnosis 页面并携带日期参数
      var dateParam = this.todayInfo;
      uni.navigateTo({
        url: "/pages/todaySymptoms/todaySymptoms?date=".concat(dateParam)
      });
    }
  }), (0, _defineProperty2.default)(_methods, "checkDiagnosisData", function checkDiagnosisData() {
    console.log('日期准备', this.todayInfo);
    if (this.dateForm.examine === false) {
      uni.showToast({
        title: '用户当日未提交检测数据',
        icon: 'none',
        duration: 2000
      });
    } else {
      // 如果 examine 为 true，跳转到 diagnosis 页面并携带日期参数
      var dateParam = this.todayInfo;
      uni.navigateTo({
        url: "/pages/experimentalInput/experimentalInput?date=".concat(dateParam)
      });
    }
  }), (0, _defineProperty2.default)(_methods, "handleFileChange", function handleFileChange(event, reportType) {
    // Handle file selection
    console.log(reportType, event);
    // Update formData based on selected files
    this.formData[reportType] = event.files;
  }), _methods)
};
exports.default = _default;
/* WEBPACK VAR INJECTION */}.call(this, __webpack_require__(/*! ./node_modules/@dcloudio/uni-mp-weixin/dist/wx.js */ 1)["default"], __webpack_require__(/*! ./node_modules/@dcloudio/uni-mp-weixin/dist/index.js */ 2)["default"]))

/***/ }),

/***/ 139:
/*!**********************************************************************************************************************************************!*\
  !*** G:/vue-init/uniapp-template-init/Infectious-disease-identification/pages/clockRecord/clockRecord.vue?vue&type=style&index=0&lang=scss& ***!
  \**********************************************************************************************************************************************/
/*! no static exports found */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_mini_css_extract_plugin_dist_loader_js_ref_8_oneOf_1_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_css_loader_dist_cjs_js_ref_8_oneOf_1_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_stylePostLoader_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_8_oneOf_1_2_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_postcss_loader_src_index_js_ref_8_oneOf_1_3_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_sass_loader_dist_cjs_js_ref_8_oneOf_1_4_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_8_oneOf_1_5_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_clockRecord_vue_vue_type_style_index_0_lang_scss___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!./node_modules/mini-css-extract-plugin/dist/loader.js??ref--8-oneOf-1-0!./node_modules/css-loader/dist/cjs.js??ref--8-oneOf-1-1!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-preprocess-loader??ref--8-oneOf-1-2!./node_modules/postcss-loader/src??ref--8-oneOf-1-3!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/sass-loader/dist/cjs.js??ref--8-oneOf-1-4!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-preprocess-loader??ref--8-oneOf-1-5!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib??vue-loader-options!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/style.js!./clockRecord.vue?vue&type=style&index=0&lang=scss& */ 140);
/* harmony import */ var _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_mini_css_extract_plugin_dist_loader_js_ref_8_oneOf_1_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_css_loader_dist_cjs_js_ref_8_oneOf_1_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_stylePostLoader_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_8_oneOf_1_2_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_postcss_loader_src_index_js_ref_8_oneOf_1_3_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_sass_loader_dist_cjs_js_ref_8_oneOf_1_4_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_8_oneOf_1_5_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_clockRecord_vue_vue_type_style_index_0_lang_scss___WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_mini_css_extract_plugin_dist_loader_js_ref_8_oneOf_1_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_css_loader_dist_cjs_js_ref_8_oneOf_1_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_stylePostLoader_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_8_oneOf_1_2_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_postcss_loader_src_index_js_ref_8_oneOf_1_3_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_sass_loader_dist_cjs_js_ref_8_oneOf_1_4_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_8_oneOf_1_5_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_clockRecord_vue_vue_type_style_index_0_lang_scss___WEBPACK_IMPORTED_MODULE_0__);
/* harmony reexport (unknown) */ for(var __WEBPACK_IMPORT_KEY__ in _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_mini_css_extract_plugin_dist_loader_js_ref_8_oneOf_1_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_css_loader_dist_cjs_js_ref_8_oneOf_1_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_stylePostLoader_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_8_oneOf_1_2_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_postcss_loader_src_index_js_ref_8_oneOf_1_3_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_sass_loader_dist_cjs_js_ref_8_oneOf_1_4_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_8_oneOf_1_5_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_clockRecord_vue_vue_type_style_index_0_lang_scss___WEBPACK_IMPORTED_MODULE_0__) if(["default"].indexOf(__WEBPACK_IMPORT_KEY__) < 0) (function(key) { __webpack_require__.d(__webpack_exports__, key, function() { return _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_mini_css_extract_plugin_dist_loader_js_ref_8_oneOf_1_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_css_loader_dist_cjs_js_ref_8_oneOf_1_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_stylePostLoader_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_8_oneOf_1_2_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_postcss_loader_src_index_js_ref_8_oneOf_1_3_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_sass_loader_dist_cjs_js_ref_8_oneOf_1_4_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_8_oneOf_1_5_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_clockRecord_vue_vue_type_style_index_0_lang_scss___WEBPACK_IMPORTED_MODULE_0__[key]; }) }(__WEBPACK_IMPORT_KEY__));
 /* harmony default export */ __webpack_exports__["default"] = (_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_mini_css_extract_plugin_dist_loader_js_ref_8_oneOf_1_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_css_loader_dist_cjs_js_ref_8_oneOf_1_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_stylePostLoader_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_8_oneOf_1_2_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_postcss_loader_src_index_js_ref_8_oneOf_1_3_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_sass_loader_dist_cjs_js_ref_8_oneOf_1_4_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_8_oneOf_1_5_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_clockRecord_vue_vue_type_style_index_0_lang_scss___WEBPACK_IMPORTED_MODULE_0___default.a); 

/***/ }),

/***/ 140:
/*!**************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js??ref--8-oneOf-1-0!./node_modules/css-loader/dist/cjs.js??ref--8-oneOf-1-1!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-preprocess-loader??ref--8-oneOf-1-2!./node_modules/postcss-loader/src??ref--8-oneOf-1-3!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/sass-loader/dist/cjs.js??ref--8-oneOf-1-4!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-preprocess-loader??ref--8-oneOf-1-5!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib??vue-loader-options!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/style.js!G:/vue-init/uniapp-template-init/Infectious-disease-identification/pages/clockRecord/clockRecord.vue?vue&type=style&index=0&lang=scss& ***!
  \**************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

// extracted by mini-css-extract-plugin
    if(false) { var cssReload; }
  

/***/ })

},[[133,"common/runtime","common/vendor"]]]);
//# sourceMappingURL=../../../.sourcemap/mp-weixin/pages/clockRecord/clockRecord.js.map