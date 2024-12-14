(global["webpackJsonp"] = global["webpackJsonp"] || []).push([["pages/todaySymptoms/todaySymptoms"],{

/***/ 165:
/*!***********************************************************************************************************************************!*\
  !*** G:/vue-init/uniapp-template-init/Infectious-disease-identification/main.js?{"page":"pages%2FtodaySymptoms%2FtodaySymptoms"} ***!
  \***********************************************************************************************************************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";
/* WEBPACK VAR INJECTION */(function(wx, createPage) {

var _interopRequireDefault = __webpack_require__(/*! @babel/runtime/helpers/interopRequireDefault */ 4);
__webpack_require__(/*! uni-pages */ 26);
var _vue = _interopRequireDefault(__webpack_require__(/*! vue */ 25));
var _todaySymptoms = _interopRequireDefault(__webpack_require__(/*! ./pages/todaySymptoms/todaySymptoms.vue */ 166));
// @ts-ignore
wx.__webpack_require_UNI_MP_PLUGIN__ = __webpack_require__;
createPage(_todaySymptoms.default);
/* WEBPACK VAR INJECTION */}.call(this, __webpack_require__(/*! ./node_modules/@dcloudio/uni-mp-weixin/dist/wx.js */ 1)["default"], __webpack_require__(/*! ./node_modules/@dcloudio/uni-mp-weixin/dist/index.js */ 2)["createPage"]))

/***/ }),

/***/ 166:
/*!****************************************************************************************************************!*\
  !*** G:/vue-init/uniapp-template-init/Infectious-disease-identification/pages/todaySymptoms/todaySymptoms.vue ***!
  \****************************************************************************************************************/
/*! no static exports found */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _todaySymptoms_vue_vue_type_template_id_b944c1c4___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./todaySymptoms.vue?vue&type=template&id=b944c1c4& */ 167);
/* harmony import */ var _todaySymptoms_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./todaySymptoms.vue?vue&type=script&lang=js& */ 169);
/* harmony reexport (unknown) */ for(var __WEBPACK_IMPORT_KEY__ in _todaySymptoms_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__) if(["default"].indexOf(__WEBPACK_IMPORT_KEY__) < 0) (function(key) { __webpack_require__.d(__webpack_exports__, key, function() { return _todaySymptoms_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__[key]; }) }(__WEBPACK_IMPORT_KEY__));
/* harmony import */ var _todaySymptoms_vue_vue_type_style_index_0_lang_css___WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./todaySymptoms.vue?vue&type=style&index=0&lang=css& */ 171);
/* harmony import */ var _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib/runtime/componentNormalizer.js */ 32);

var renderjs





/* normalize component */

var component = Object(_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _todaySymptoms_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__["default"],
  _todaySymptoms_vue_vue_type_template_id_b944c1c4___WEBPACK_IMPORTED_MODULE_0__["render"],
  _todaySymptoms_vue_vue_type_template_id_b944c1c4___WEBPACK_IMPORTED_MODULE_0__["staticRenderFns"],
  false,
  null,
  null,
  null,
  false,
  _todaySymptoms_vue_vue_type_template_id_b944c1c4___WEBPACK_IMPORTED_MODULE_0__["components"],
  renderjs
)

component.options.__file = "pages/todaySymptoms/todaySymptoms.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ 167:
/*!***********************************************************************************************************************************************!*\
  !*** G:/vue-init/uniapp-template-init/Infectious-disease-identification/pages/todaySymptoms/todaySymptoms.vue?vue&type=template&id=b944c1c4& ***!
  \***********************************************************************************************************************************************/
/*! exports provided: render, staticRenderFns, recyclableRender, components */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_17_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_template_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_uni_app_loader_page_meta_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_todaySymptoms_vue_vue_type_template_id_b944c1c4___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-preprocess-loader??ref--17-0!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/template.js!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-uni-app-loader/page-meta.js!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib??vue-loader-options!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/style.js!./todaySymptoms.vue?vue&type=template&id=b944c1c4& */ 168);
/* harmony reexport (safe) */ __webpack_require__.d(__webpack_exports__, "render", function() { return _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_17_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_template_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_uni_app_loader_page_meta_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_todaySymptoms_vue_vue_type_template_id_b944c1c4___WEBPACK_IMPORTED_MODULE_0__["render"]; });

/* harmony reexport (safe) */ __webpack_require__.d(__webpack_exports__, "staticRenderFns", function() { return _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_17_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_template_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_uni_app_loader_page_meta_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_todaySymptoms_vue_vue_type_template_id_b944c1c4___WEBPACK_IMPORTED_MODULE_0__["staticRenderFns"]; });

/* harmony reexport (safe) */ __webpack_require__.d(__webpack_exports__, "recyclableRender", function() { return _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_17_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_template_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_uni_app_loader_page_meta_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_todaySymptoms_vue_vue_type_template_id_b944c1c4___WEBPACK_IMPORTED_MODULE_0__["recyclableRender"]; });

/* harmony reexport (safe) */ __webpack_require__.d(__webpack_exports__, "components", function() { return _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_17_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_template_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_uni_app_loader_page_meta_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_todaySymptoms_vue_vue_type_template_id_b944c1c4___WEBPACK_IMPORTED_MODULE_0__["components"]; });



/***/ }),

/***/ 168:
/*!***********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-preprocess-loader??ref--17-0!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/template.js!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-uni-app-loader/page-meta.js!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib??vue-loader-options!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/style.js!G:/vue-init/uniapp-template-init/Infectious-disease-identification/pages/todaySymptoms/todaySymptoms.vue?vue&type=template&id=b944c1c4& ***!
  \***********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
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
  if (!_vm._isMounted) {
    _vm.e0 = function ($event) {
      _vm.temporaryChoice1 = !_vm.temporaryChoice1
    }
    _vm.e1 = function ($event) {
      _vm.temporaryChoice2 = !_vm.temporaryChoice2
    }
    _vm.e2 = function ($event) {
      _vm.otherChoices = !_vm.otherChoices
    }
    _vm.e3 = function ($event) {
      _vm.SeafoodChoices = !_vm.SeafoodChoices
    }
  }
}
var recyclableRender = false
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ 169:
/*!*****************************************************************************************************************************************!*\
  !*** G:/vue-init/uniapp-template-init/Infectious-disease-identification/pages/todaySymptoms/todaySymptoms.vue?vue&type=script&lang=js& ***!
  \*****************************************************************************************************************************************/
/*! no static exports found */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_babel_loader_lib_index_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_13_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_script_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_todaySymptoms_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!./node_modules/babel-loader/lib!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-preprocess-loader??ref--13-1!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/script.js!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib??vue-loader-options!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/style.js!./todaySymptoms.vue?vue&type=script&lang=js& */ 170);
/* harmony import */ var _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_babel_loader_lib_index_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_13_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_script_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_todaySymptoms_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_babel_loader_lib_index_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_13_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_script_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_todaySymptoms_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__);
/* harmony reexport (unknown) */ for(var __WEBPACK_IMPORT_KEY__ in _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_babel_loader_lib_index_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_13_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_script_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_todaySymptoms_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__) if(["default"].indexOf(__WEBPACK_IMPORT_KEY__) < 0) (function(key) { __webpack_require__.d(__webpack_exports__, key, function() { return _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_babel_loader_lib_index_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_13_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_script_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_todaySymptoms_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__[key]; }) }(__WEBPACK_IMPORT_KEY__));
 /* harmony default export */ __webpack_exports__["default"] = (_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_babel_loader_lib_index_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_13_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_script_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_todaySymptoms_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0___default.a); 

/***/ }),

/***/ 170:
/*!************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-preprocess-loader??ref--13-1!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/script.js!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib??vue-loader-options!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/style.js!G:/vue-init/uniapp-template-init/Infectious-disease-identification/pages/todaySymptoms/todaySymptoms.vue?vue&type=script&lang=js& ***!
  \************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";
/* WEBPACK VAR INJECTION */(function(uni) {

var _interopRequireDefault = __webpack_require__(/*! @babel/runtime/helpers/interopRequireDefault */ 4);
Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = void 0;
var _regenerator = _interopRequireDefault(__webpack_require__(/*! @babel/runtime/regenerator */ 42));
var _asyncToGenerator2 = _interopRequireDefault(__webpack_require__(/*! @babel/runtime/helpers/asyncToGenerator */ 44));
var _index = _interopRequireDefault(__webpack_require__(/*! ../../store/index.ts */ 63));
var _clock = __webpack_require__(/*! ../../services/clock */ 90);
//
var _default = {
  data: function data() {
    return {
      AllCheckInfos: {},
      AllCheckInfo: {
        circulatorysymptoms: {
          circulatorySymptomsId: 0,
          hasArrhythmia: false,
          hasChestPain: false,
          hasLowBloodPressure: false,
          hasPalpitation: false,
          hasRapidPulse: false,
          statusId: 0
        },
        digestivesymptoms: {
          arrheaFrequencyGEThreeTimesPerDay: false,
          diarrheaMode: null,
          diarrheaTimesPerDay: '',
          digestiveSymptomsId: 0,
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
          statusId: 0,
          stoolAmount: null,
          stoolOdor: null,
          stoolType1Detail: null,
          stoolType2Detail: null,
          vomitingMode: '先泻后吐'
        },
        generalsymptoms: {
          dehydration: null,
          feverDuration: '',
          generalSymptomsId: 0,
          hasBloodBlisters: false,
          hasChills: false,
          hasCongestiveOrPetechialRash: false,
          hasCyanosis: false,
          hasDehydration: false,
          hasEdema: false,
          hasExhaustion: false,
          hasFatigue: false,
          hasFever: false,
          hasHeadache: false,
          hasItchyRash: false,
          hasJointPain: false,
          hasLymphNodeSwelling: false,
          hasLymphNodeSwellingArmpit: false,
          hasLymphNodeSwellingElbow: false,
          hasLymphNodeSwellingGroin: false,
          hasLymphNodeSwellingNeck: false,
          hasLymphNodeSwellingPopliteal: false,
          hasLymphNodeSwellingSubclavian: false,
          hasMusclePain: false,
          hasNightSweats: false,
          hasPainfulRedRash: false,
          hasSkinUlcer: false,
          hasSubcutaneousAndMucosalBleedingSpots: false,
          hasSweating: false,
          hasWeightLoss: false,
          highestTemperature: '',
          itchyRashOnBackOfHands: false,
          itchyRashOnFace: false,
          itchyRashOnFeet: false,
          itchyRashOnFingers: false,
          itchyRashOnLowerLimbs: false,
          itchyRashOnOther: false,
          itchyRashOnUpperLimbs: false,
          statusId: 0
        },
        localSymptoms: {
          hasConjunctivalCongestion: false,
          hasConjunctivitis: false,
          hasEyelidEdema: false,
          hasFacialRednessOrPallor: false,
          hasGumBleeding: false,
          hasKidneyPain: false,
          hasLipCyanosis: false,
          hasLowBackPain: false,
          hasMucosalBleedingInOralOrNasal: false,
          hasOrbitalPain: false,
          hasRednessNeckChest: false,
          hasSkinBleedingPoints: false,
          hasSkinPallorOrBruising: false,
          hasSmellTasteLoss: false,
          localSymptomsId: 0,
          statusId: 0
        },
        neurologicalSymptoms: {
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
          hasTinnitus: false,
          neurologicalSymptomsId: 0,
          statusId: 0
        },
        otherSymptoms: {
          hasCalfMusclePain: false,
          hasForcedPosture: false,
          hasOtherSymptoms: false,
          hasPeriodicAttack: false,
          hasRapidProgress: false,
          hasSuddenOnset: false,
          nutritionStatus: '好',
          otherSymptomsId: 0,
          otherSymptomsName: '',
          sleepQuality: '好',
          statusId: 0,
          symptomSeverity: '轻',
          workLifeStress: '小'
        },
        respiratorysymptoms: {
          hasRunnyNose: false,
          hasBloodySputum: false,
          hasBreathingDifficulty: false,
          hasBronchitis: false,
          hasChestTightness: false,
          hasCough: false,
          hasDryThroat: false,
          hasHemoptysis: false,
          hasHoarseness: false,
          hasNasalCongestion: false,
          hasShortnessOfBreath: false,
          hasSoreThroat: false,
          hasSputum: false,
          respiratorySymptomsId: 0,
          statusId: 0
        },
        riskFactorsAndExposure: {
          anthraxArea: '否',
          contactDate: '',
          contactPlace: '',
          contactWithAnimalProducts: '否',
          contactWithAnimals: '否',
          contactWithDiarrheaPatient: '否',
          contactWithFeverPatient: '否',
          contactWithPatientExcreta: '否',
          contactWithRashPatient: '否',
          contactWithRatIn2Months: '否',
          contactWithTuberculosisPatient: '否',
          diningOut: false,
          diningPlace: '',
          diningUnclear: false,
          drunkWaterFromDitchesOrPonds: '否',
          eatenFoodContaminatedByRatFeces: '否',
          fleaBite: '否',
          foodNames: '',
          groupDining: false,
          groupOutbreak: '否',
          hasContactWithBird: false,
          hasContactWithCat: false,
          hasContactWithCow: false,
          hasContactWithDog: false,
          hasContactWithDonkey: false,
          hasContactWithFlea: false,
          hasContactWithHorse: false,
          hasContactWithLice: false,
          hasContactWithMarmot: false,
          hasContactWithMule: false,
          hasContactWithOtherAnimal: false,
          hasContactWithPig: false,
          hasContactWithRabbit: false,
          hasContactWithRat: false,
          hasContactWithSheep: false,
          hasContactWithTiger: false,
          hasContactWithWolf: false,
          hasDrinkingRawWater: false,
          hasEatingColdCookedFood: false,
          hasEatingRawFood: false,
          hasEatingSeafood: false,
          hasMosquitoBite: false,
          hillyOrMountainous: false,
          malariaArea: '否',
          otherAnimalName: '',
          otherResidenceType: false,
          otherWildActivityName: '',
          outdoorStayOrWorkWithinMonth: '否',
          plagueArea: '否',
          plain: false,
          ratContactMethod: null,
          ratOrRatDroppingsAtWorkplace: '否',
          restedNearRatHoles: '否',
          riskFactorsAndExposureId: 0,
          specificAnthraxLocation: '',
          specificMalariaLocation: '',
          specificPlagueLocation: '',
          statusId: 0,
          tickBite: '否',
          wildActivityHistory: '其他'
        }
      },
      formData: {
        option: '' // 用于存储选中的选项
      },

      temporaryChoice1: false,
      temporaryChoice2: false,
      otherChoices: false,
      SeafoodChoices: false,
      isViewMode: false,
      pages: 1,
      sonPage: 1
    };
  },
  onLoad: function onLoad(options) {
    // `options` 是页面加载时获取的传参对象
    var date = options.date;
    console.log('dedao', date, '那么isViewMode', this.isViewMode);
    if (date) {
      this.isViewMode = true;
      // 调用获取服务方法以查询数据
      this.getDiagnosisData(date);
    }
  },
  methods: {
    getDiagnosisData: function getDiagnosisData(date) {
      var _this = this;
      return (0, _asyncToGenerator2.default)( /*#__PURE__*/_regenerator.default.mark(function _callee() {
        var response;
        return _regenerator.default.wrap(function _callee$(_context) {
          while (1) {
            switch (_context.prev = _context.next) {
              case 0:
                _context.prev = 0;
                _context.next = 3;
                return (0, _clock.getDailyReportService)(date);
              case 3:
                response = _context.sent;
                if (response && response.data) {
                  _this.AllCheckInfo = response.data;
                }
                _context.next = 10;
                break;
              case 7:
                _context.prev = 7;
                _context.t0 = _context["catch"](0);
                console.error('获取诊断数据失败:', _context.t0);
              case 10:
              case "end":
                return _context.stop();
            }
          }
        }, _callee, null, [[0, 7]]);
      }))();
    },
    toggleSymptom: function toggleSymptom(symptom) {
      if (this.symptoms.includes(symptom)) {
        this.symptoms = this.symptoms.filter(function (s) {
          return s !== symptom;
        });
      } else {
        this.symptoms.push(symptom);
      }
    },
    handleSubmit: function handleSubmit() {
      // 处理提交逻辑
      console.log('提交数据', {
        fever: this.fever,
        duration: this.selectedDuration,
        symptoms: this.symptoms
      });
      if (this.pages < 8) {
        this.pages++;
      } else {
        if (this.sonPage <= 5) {
          this.sonPage++;
        }
      }

      // uni.navigateTo({
      // 	url: '/pages/respiratorySymptoms/respiratorySymptoms'
      // })
    },
    showRadio: function showRadio(maleValue) {
      this.AllCheckInfo.generalsymptoms.dehydration = maleValue;
      console.log('已经选择', this.AllCheckInfo.generalsymptoms.dehydration);
    },
    handleCheckboxChange: function handleCheckboxChange(propName) {
      console.log("Changing ".concat(propName, ":"), this.AllCheckInfo.generalsymptoms[propName]);
      this.AllCheckInfo.generalsymptoms[propName] = !this.AllCheckInfo.generalsymptoms[propName];
    },
    handleRespiratorysymptomsBox: function handleRespiratorysymptomsBox(propName) {
      console.log("Changing ".concat(propName, ":"), this.AllCheckInfo.respiratorysymptoms[propName]);
      this.AllCheckInfo.respiratorysymptoms[propName] = !this.AllCheckInfo.respiratorysymptoms[propName];
    },
    handlereturn: function handlereturn() {
      if (this.pages == 8 && this.sonPage > 1) {
        this.sonPage--;
      } else {
        this.pages--;
      }
    },
    submitSuccess1: function submitSuccess1(response) {
      console.log(response);
    },
    handleFinallySubmit: function handleFinallySubmit() {
      var _this2 = this;
      uni.showModal({
        title: '提示',
        content: '您确定要提交吗？',
        confirmText: '确定',
        cancelText: '取消',
        success: function success(res) {
          if (res.confirm) {
            var that = _this2; // 保存当前的 `this` 上下文
            var clockForm = _index.default.state.clockFormParams;
            // 使用顺序执行的方式确保顺序执行
            (0, _clock.clockSubmitService)(clockForm).then(function (res) {
              _this2.submitSuccess1(res);

              // 在 `clockSubmitService` 完成后再执行 `submitAllCheckService`
              return (0, _clock.submitAllCheckService)(that.AllCheckInfo);
            }).then(function (res) {
              that.submitSuccess(res);
            }).catch(function (error) {
              that.submitError(error);
            });
          }
        }
      });
    },
    submitError: function submitError(error) {
      uni.showToast({
        title: '信息提交失败',
        icon: 'none',
        // 可选，默认为none，还有success, loading, warning, error等
        duration: 1500 // 显示持续时间，单位为毫秒，默认为1500
      });
    },
    submitSuccess: function submitSuccess(response) {
      console.log(response);
      console.log(response.data);
      _index.default.dispatch('setDiseasescoring', response.data);
      if (response.code == 1) {
        uni.showToast({
          title: '信息提交成功',
          icon: 'success',
          duration: 1800
        });
        setTimeout(function () {
          // 进行页面跳转
          uni.navigateTo({
            url: '/pages/clockFail/clockFail'
          });
        }, 50);
      } else {
        uni.showModal({
          title: '信息提交失败',
          content: '信息提交失败。',
          showCancel: false,
          confirmText: '确定',
          success: function success(res) {}
        });
      }
    },
    handleDigestFrequency: function handleDigestFrequency(propName) {
      console.log("Changing ".concat(propName, ":"), this.AllCheckInfo.digestivesymptoms[propName]);
      this.AllCheckInfo.digestivesymptoms[propName] = !this.AllCheckInfo.digestivesymptoms[propName];
    },
    showDigestRadio: function showDigestRadio(maleValue) {
      this.AllCheckInfo.digestivesymptoms.diarrheaMode = maleValue;
      console.log('已经选择', this.AllCheckInfo.digestivesymptoms.diarrheaMode);
    },
    showDigestRadioBystool1: function showDigestRadioBystool1(maleValue) {
      this.AllCheckInfo.digestivesymptoms.stoolType1Detail = maleValue;
      console.log('已经选择', this.AllCheckInfo.digestivesymptoms.stoolType1Detail);
    },
    showDigestRadioBystool2: function showDigestRadioBystool2(maleValue) {
      this.AllCheckInfo.digestivesymptoms.stoolType2Detail = maleValue;
      console.log('已经选择', this.AllCheckInfo.digestivesymptoms.stoolType2Detail);
    },
    showDigestVomitingMode: function showDigestVomitingMode(maleValue) {
      this.AllCheckInfo.digestivesymptoms.vomitingMode = maleValue;
      console.log('已经选择', this.AllCheckInfo.digestivesymptoms.vomitingMode);
    },
    handleCirculatorysymptoms: function handleCirculatorysymptoms(propName) {
      console.log("Changing ".concat(propName, ":"), this.AllCheckInfo.circulatorysymptoms[propName]);
      this.AllCheckInfo.circulatorysymptoms[propName] = !this.AllCheckInfo.circulatorysymptoms[propName];
    },
    handleNeurologicalSymptoms: function handleNeurologicalSymptoms(propName) {
      console.log("Changing ".concat(propName, ":"), this.AllCheckInfo.neurologicalSymptoms[propName]);
      this.AllCheckInfo.neurologicalSymptoms[propName] = !this.AllCheckInfo.neurologicalSymptoms[propName];
    },
    handleLocalSymptoms: function handleLocalSymptoms(propName) {
      console.log("Changing ".concat(propName, ":"), this.AllCheckInfo.localSymptoms[propName]);
      this.AllCheckInfo.localSymptoms[propName] = !this.AllCheckInfo.localSymptoms[propName];
    },
    handleOtherSymptoms: function handleOtherSymptoms(propName) {
      console.log("Changing ".concat(propName, ":"), this.AllCheckInfo.otherSymptoms[propName]);
      this.AllCheckInfo.otherSymptoms[propName] = !this.AllCheckInfo.otherSymptoms[propName];
    },
    showSleepQuality: function showSleepQuality(maleValue) {
      this.AllCheckInfo.otherSymptoms.sleepQuality = maleValue;
      console.log('已经选择', this.AllCheckInfo.otherSymptoms.sleepQuality);
    },
    showNutritionStatus: function showNutritionStatus(maleValue) {
      this.AllCheckInfo.otherSymptoms.nutritionStatus = maleValue;
      console.log('已经选择', this.AllCheckInfo.otherSymptoms.nutritionStatus);
    },
    showWorkLifeStress: function showWorkLifeStress(maleValue) {
      this.AllCheckInfo.otherSymptoms.workLifeStress = maleValue;
      console.log('已经选择', this.AllCheckInfo.otherSymptoms.workLifeStress);
    },
    showSymptomSeverity: function showSymptomSeverity(maleValue) {
      this.AllCheckInfo.otherSymptoms.symptomSeverity = maleValue;
      console.log('已经选择', this.AllCheckInfo.otherSymptoms.symptomSeverity);
    },
    shwoContactWithFeverPatient: function shwoContactWithFeverPatient(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.contactWithFeverPatient = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.contactWithFeverPatient);
    },
    showContactWithDiarrheaPatient: function showContactWithDiarrheaPatient(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.contactWithDiarrheaPatient = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.contactWithDiarrheaPatient);
    },
    showContactWithRashPatient: function showContactWithRashPatient(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.contactWithRashPatient = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.contactWithRashPatient);
    },
    showContactWithTuberculosisPatient: function showContactWithTuberculosisPatient(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.contactWithTuberculosisPatient = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.contactWithTuberculosisPatient);
    },
    handleFactorSymptoms: function handleFactorSymptoms(propName) {
      console.log("Changing ".concat(propName, ":"), this.AllCheckInfo.riskFactorsAndExposure[propName]);
      this.AllCheckInfo.riskFactorsAndExposure[propName] = !this.AllCheckInfo.riskFactorsAndExposure[propName];
    },
    showGroupOutbreak: function showGroupOutbreak(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.groupOutbreak = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.groupOutbreak);
    },
    showOutdoorStayOrWorkWithinMonth: function showOutdoorStayOrWorkWithinMonth(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.outdoorStayOrWorkWithinMonth = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.outdoorStayOrWorkWithinMonth);
    },
    showWildActivityHistory: function showWildActivityHistory(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.wildActivityHistory = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.wildActivityHistory);
    },
    showPlagueArea: function showPlagueArea(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.plagueArea = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.plagueArea);
    },
    showAnthraxArea: function showAnthraxArea(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.anthraxArea = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.anthraxArea);
    },
    showMalariaArea: function showMalariaArea(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.malariaArea = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.malariaArea);
    },
    showTickBite: function showTickBite(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.tickBite = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.tickBite);
    },
    showFleaBite: function showFleaBite(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.fleaBite = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.fleaBite);
    },
    showContactWithRatIn2Months: function showContactWithRatIn2Months(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.contactWithRatIn2Months = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.contactWithRatIn2Months);
      if (this.AllCheckInfo.riskFactorsAndExposure.contactWithRatIn2Months == '否') {
        this.AllCheckInfo.riskFactorsAndExposure.ratContactMethod = null;
      }
    },
    showRatContactMethod: function showRatContactMethod(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.ratContactMethod = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.ratContactMethod);
    },
    showEatenFoodContaminatedByRatFeces: function showEatenFoodContaminatedByRatFeces(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.eatenFoodContaminatedByRatFeces = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.eatenFoodContaminatedByRatFeces);
    },
    showDrunkWaterFromDitchesOrPonds: function showDrunkWaterFromDitchesOrPonds(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.drunkWaterFromDitchesOrPonds = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.drunkWaterFromDitchesOrPonds);
    },
    showRestedNearRatHoles: function showRestedNearRatHoles(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.restedNearRatHoles = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.restedNearRatHoles);
    },
    showRatOrRatDroppingsAtWorkplace: function showRatOrRatDroppingsAtWorkplace(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.ratOrRatDroppingsAtWorkplace = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.ratOrRatDroppingsAtWorkplace);
    },
    showContactWithPatientExcreta: function showContactWithPatientExcreta(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.contactWithPatientExcreta = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.contactWithPatientExcreta);
    },
    showOtherWildActivityName: function showOtherWildActivityName(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.otherWildActivityName = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.otherWildActivityName);
    },
    showContactWithAnimals: function showContactWithAnimals(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.contactWithAnimals = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.contactWithAnimals);
      if (this.AllCheckInfo.riskFactorsAndExposure.contactWithAnimals === '否') {
        this.AllCheckInfo.riskFactorsAndExposure.hasContactWithRat = false;
        this.AllCheckInfo.riskFactorsAndExposure.hasContactWithLice = false;
        this.AllCheckInfo.riskFactorsAndExposure.hasContactWithFlea = false;
        this.AllCheckInfo.riskFactorsAndExposure.hasContactWithCat = false;
        this.AllCheckInfo.riskFactorsAndExposure.hasContactWithMarmot = false;
        this.AllCheckInfo.riskFactorsAndExposure.hasContactWithBird = false;
        this.AllCheckInfo.riskFactorsAndExposure.hasContactWithPig = false;
        this.AllCheckInfo.riskFactorsAndExposure.hasContactWithHorse = false;
        this.AllCheckInfo.riskFactorsAndExposure.hasContactWithMule = false;
        this.AllCheckInfo.riskFactorsAndExposure.hasContactWithDonkey = false;
        this.AllCheckInfo.riskFactorsAndExposure.hasContactWithCow = false;
        this.AllCheckInfo.riskFactorsAndExposure.hasContactWithSheep = false;
        this.AllCheckInfo.riskFactorsAndExposure.hasContactWithDog = false;
        this.AllCheckInfo.riskFactorsAndExposure.hasContactWithRabbit = false;
        this.AllCheckInfo.riskFactorsAndExposure.hasContactWithWolf = false;
        this.AllCheckInfo.riskFactorsAndExposure.hasContactWithTiger = false;
      }
    },
    showContactWithAnimalProducts: function showContactWithAnimalProducts(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.contactWithAnimalProducts = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.contactWithAnimalProducts);
    },
    showStoolOdor: function showStoolOdor(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.stoolOdor = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.stoolOdor);
    },
    showStoolAmount: function showStoolAmount(maleValue) {
      this.AllCheckInfo.riskFactorsAndExposure.stoolAmount = maleValue;
      console.log('已经选择', this.AllCheckInfo.riskFactorsAndExposure.stoolAmount);
    },
    onDateChange: function onDateChange(e) {
      this.AllCheckInfo.riskFactorsAndExposure.contactDate = e.detail.value;
    }
  }
};
exports.default = _default;
/* WEBPACK VAR INJECTION */}.call(this, __webpack_require__(/*! ./node_modules/@dcloudio/uni-mp-weixin/dist/index.js */ 2)["default"]))

/***/ }),

/***/ 171:
/*!*************************************************************************************************************************************************!*\
  !*** G:/vue-init/uniapp-template-init/Infectious-disease-identification/pages/todaySymptoms/todaySymptoms.vue?vue&type=style&index=0&lang=css& ***!
  \*************************************************************************************************************************************************/
/*! no static exports found */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_mini_css_extract_plugin_dist_loader_js_ref_6_oneOf_1_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_css_loader_dist_cjs_js_ref_6_oneOf_1_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_stylePostLoader_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_6_oneOf_1_2_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_postcss_loader_src_index_js_ref_6_oneOf_1_3_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_todaySymptoms_vue_vue_type_style_index_0_lang_css___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!./node_modules/mini-css-extract-plugin/dist/loader.js??ref--6-oneOf-1-0!./node_modules/css-loader/dist/cjs.js??ref--6-oneOf-1-1!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-preprocess-loader??ref--6-oneOf-1-2!./node_modules/postcss-loader/src??ref--6-oneOf-1-3!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib??vue-loader-options!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/style.js!./todaySymptoms.vue?vue&type=style&index=0&lang=css& */ 172);
/* harmony import */ var _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_mini_css_extract_plugin_dist_loader_js_ref_6_oneOf_1_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_css_loader_dist_cjs_js_ref_6_oneOf_1_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_stylePostLoader_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_6_oneOf_1_2_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_postcss_loader_src_index_js_ref_6_oneOf_1_3_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_todaySymptoms_vue_vue_type_style_index_0_lang_css___WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_mini_css_extract_plugin_dist_loader_js_ref_6_oneOf_1_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_css_loader_dist_cjs_js_ref_6_oneOf_1_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_stylePostLoader_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_6_oneOf_1_2_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_postcss_loader_src_index_js_ref_6_oneOf_1_3_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_todaySymptoms_vue_vue_type_style_index_0_lang_css___WEBPACK_IMPORTED_MODULE_0__);
/* harmony reexport (unknown) */ for(var __WEBPACK_IMPORT_KEY__ in _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_mini_css_extract_plugin_dist_loader_js_ref_6_oneOf_1_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_css_loader_dist_cjs_js_ref_6_oneOf_1_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_stylePostLoader_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_6_oneOf_1_2_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_postcss_loader_src_index_js_ref_6_oneOf_1_3_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_todaySymptoms_vue_vue_type_style_index_0_lang_css___WEBPACK_IMPORTED_MODULE_0__) if(["default"].indexOf(__WEBPACK_IMPORT_KEY__) < 0) (function(key) { __webpack_require__.d(__webpack_exports__, key, function() { return _D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_mini_css_extract_plugin_dist_loader_js_ref_6_oneOf_1_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_css_loader_dist_cjs_js_ref_6_oneOf_1_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_stylePostLoader_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_6_oneOf_1_2_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_postcss_loader_src_index_js_ref_6_oneOf_1_3_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_todaySymptoms_vue_vue_type_style_index_0_lang_css___WEBPACK_IMPORTED_MODULE_0__[key]; }) }(__WEBPACK_IMPORT_KEY__));
 /* harmony default export */ __webpack_exports__["default"] = (_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_mini_css_extract_plugin_dist_loader_js_ref_6_oneOf_1_0_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_css_loader_dist_cjs_js_ref_6_oneOf_1_1_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_loaders_stylePostLoader_js_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_webpack_preprocess_loader_index_js_ref_6_oneOf_1_2_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_postcss_loader_src_index_js_ref_6_oneOf_1_3_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_vue_cli_plugin_uni_packages_vue_loader_lib_index_js_vue_loader_options_D_cxdownload_HBuilderX_4_24_2024072208_HBuilderX_plugins_uniapp_cli_node_modules_dcloudio_webpack_uni_mp_loader_lib_style_js_todaySymptoms_vue_vue_type_style_index_0_lang_css___WEBPACK_IMPORTED_MODULE_0___default.a); 

/***/ }),

/***/ 172:
/*!*****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js??ref--6-oneOf-1-0!./node_modules/css-loader/dist/cjs.js??ref--6-oneOf-1-1!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/webpack-preprocess-loader??ref--6-oneOf-1-2!./node_modules/postcss-loader/src??ref--6-oneOf-1-3!./node_modules/@dcloudio/vue-cli-plugin-uni/packages/vue-loader/lib??vue-loader-options!./node_modules/@dcloudio/webpack-uni-mp-loader/lib/style.js!G:/vue-init/uniapp-template-init/Infectious-disease-identification/pages/todaySymptoms/todaySymptoms.vue?vue&type=style&index=0&lang=css& ***!
  \*****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

// extracted by mini-css-extract-plugin
    if(false) { var cssReload; }
  

/***/ })

},[[165,"common/runtime","common/vendor"]]]);
//# sourceMappingURL=../../../.sourcemap/mp-weixin/pages/todaySymptoms/todaySymptoms.js.map