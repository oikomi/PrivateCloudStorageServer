/* 
 * @Author: White
 * @Email: weifengwang@pptv.com
 * @Date:   2015-03-29 20:12:03
 * @Last Modified time: 2015-04-13 20:59:33
 */
define(function (require, exports, module) {
    var url1 = './api/v1/storage';
    var url2 = './open';
    var api = {
        login: url1 + '?action=login',
        getFileList: url1 + '?action=get_server_file_list',
        mkdir: url1 + '?action=mkdir',
        rmfile: url1 + '?action=rm_file',
        rename: url1 + '?action=rename',
        upload: url1 + '?action=upload',
        download: url2,
        moveFile:url1+'?action=move_file'
        // login: './login.php',
        // getFileList: './getFileList.php',
        // mkdir: './mkdir.php',
        // rmfile: './mkdir.php',
        // rename: './mkdir.php',
        // upload: 'http://54.187.96.250:10000/api/v1/storage?action=upload',
        // moveFile:'./mkdir.php'
    };
    return api;
})