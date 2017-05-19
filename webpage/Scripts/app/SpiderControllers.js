"use strict";

/* controllers module */

angular.module("BsTableDirective.SpiderControllers", ["BsTableDirective.Services","ui.bootstrap"])
    .controller("SpiderCtrl", ["$scope", "BootswatchService","$http", 
        function ($scope, BootswatchService, $http) {
        // scope models
        $scope.page = { Css: "/statics/bootstrap/css/bootstrap.min.css" };
        $scope.contactList = [];
        $scope.progress = { Ready: false };
        $scope.showspider = "";
        $scope.showid = 0;
        // get themes from bootswatch
        BootswatchService.GetAll().success(function (result) {
            // set themes to scope
            $scope.page.Themes = result.themes;
        });
        // get data for bs-table
        GenerateData();
        $scope.TestSpider = function() {

        }
        $scope.SaveSpider = function(){

        }
        // generate data
        function GenerateData() {
            $http({
                method: 'GET',
                url: '/GetSpiderConfigJson'
            }).then(function successCallback(response) {
                // 请求成功执行代码

                $scope.contactList = response.data.ret;
                console.log(' $scope.contactList = ', $scope.contactList); 
                $scope.showspider = $scope.contactList[$scope.showid];
                // hide progress
                $scope.progress = { Ready: true };
                //$scope.$apply();
            }, function errorCallback(response) {
            // 请求失败执行代码
                alert("Http Get Error:GetTNovelTableInfo");

            });
        }
        // init socialite jquery plugin
        Socialite.load();
    }]);
    