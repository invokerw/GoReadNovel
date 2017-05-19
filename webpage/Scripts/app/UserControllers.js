"use strict";

/* controllers module */

angular.module("BsTableDirective.UserControllers", ["BsTableDirective.Services","ui.bootstrap"])
    .controller("UserCtrl", ["$scope", "BootswatchService","$http", 
        function ($scope, BootswatchService, $http) {
        // scope models
        $scope.page = { Css: "/statics/bootstrap/css/bootstrap.min.css" };
        $scope.contactList = [];
        $scope.progress = { Ready: false };
        // get themes from bootswatch
        BootswatchService.GetAll().success(function (result) {
            // set themes to scope
            $scope.page.Themes = result.themes;
            // hide progress
            //$scope.progress = { Ready: true };
        });
        // get data for bs-table
        GenerateData();
        // show function for bs-table
        $scope.Show = function (contact) {
            alert(JSON.stringify(contact));
        };
        // edit function for bs-table
        $scope.Edit = function (contact) {
            alert(JSON.stringify(contact));
        };
        // remove function for bs-table
        $scope.Remove = function (contact) {
            alert(JSON.stringify(contact));
        };
        // generate data
        function GenerateData() {
            $http({
                method: 'GET',
                url: '/GetUsersInfoJson'
            }).then(function successCallback(response) {
                // 请求成功执行代码
                console.log('response.data = ',response.data);   
                // hide progress
                $scope.progress = { Ready: true };
                $scope.contactList = response.data.ret;
                //return dataList;
            }, function errorCallback(response) {
            // 请求失败执行代码
                alert("Http Get Error:GetTNovelTableInfo");

            });
        }
        // init socialite jquery plugin
        Socialite.load();
    }]);
    