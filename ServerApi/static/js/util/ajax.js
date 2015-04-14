/* 
* @Author: White
* @Email: weifengwang@pptv.com
* @Date:   2015-03-29 19:27:52
* @Last Modified time: 2015-03-30 23:30:51
*/
define(function (require, exports, module){
    var $ = require('jquery');
    function ajax(option){
        var opt = $.extend(true, {
            type: 'GET',
            dataType: 'json',
            success: function(){},
            error: function(){}
        }, option);
        $.ajax(opt);
    }
    return ajax;
});