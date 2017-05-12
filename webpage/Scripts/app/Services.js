"use strict";

/* services module */

//var bootsWatchApi = "http://bootswatch.com/api/3.json";
var bootsWatchApi = "/statics/json/3.json";
angular.module("BsTableDirective.Services", [])
    .factory("BootswatchService", ["$http", function ($http) {
        return {
            GetAll: function () {
                return $http.get(bootsWatchApi);
            }
        }
    }]);
