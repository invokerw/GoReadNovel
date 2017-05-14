﻿"use strict";

/* app module */

angular.module("BsTableDirective", ["ngRoute", "bsTable", "BsTableDirective.Controllers", 
	"BsTableDirective.Services"])
    .config(["$routeProvider", function ($routeProvider) {
        $routeProvider.when("/", {
            templateUrl: "/Content/templates/example.html",
            controller: "ExampleCtrl"
        })
        .when("/ShowNovelInfo", {
            templateUrl: "/Content/templates/noveltable.html",
            controller: "NovelInfoCtrl"
        })
        .otherwise({ redirectTo: "/" });
    }])
    .filter('removal',function(){
    	return function(text) {
    	//var str = "http://www.huanyue123.com/book"; 
    	//str.length = 30;
        return text.substring(30);
    	}
    });
