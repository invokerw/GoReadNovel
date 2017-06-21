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
        $scope.oksave = false;
        $scope.constantlength = 0;
        // get themes from bootswatch
        BootswatchService.GetAll().success(function (result) {
            // set themes to scope
            $scope.page.Themes = result.themes;
            $scope.page.Css = $scope.page.Themes[15].cssMin;
        });
        // get data for bs-table
        GenerateData();
        $scope.selectMe = function(index){
            $scope.showid = index;
            $scope.oksave = false;
            $scope.showspider = $scope.contactList[index];
        }
        $scope.textchange = function(textvalue,index){
            console.log(textvalue,index);
            $scope.contactList[$scope.showid].textvalues[index] = textvalue;
        }
        $scope.testchange = function(testvalue,index){
            console.log(testvalue,index);
            $scope.contactList[$scope.showid].testvalues[index] = testvalue;
        }
        $scope.TestSpider = function() {
            //alert(JSON.stringify($scope.contactList[$scope.showid]))
            $http({
                method: 'GET',
                url: '/TestConfigJson',
                params: {
                    config: $scope.contactList[$scope.showid]
                    },
            }).then(function successCallback(response) {
                // 请求成功执行代码
                console.log('response.data= ', response.data); 
                alert(JSON.stringify(response.data))
                if (response.data.code == 1) {
                    $scope.showspider = $scope.contactList[$scope.showid];
                    $scope.oksave = true;
                }

            }, function errorCallback(response) {
            // 请求失败执行代码
                alert("Http Error:TestConfigJson");

            });
        }
        $scope.SaveSpider = function(){
            if ($scope.oksave == false) return;  
            $http({
                method: 'GET',
                url: '/SaveConfigJson',

            }).then(function successCallback(response) {
                // 请求成功执行代码
                console.log('response.data= ', response.data); 
                if (response.data.code == 1) {
                    $scope.oksave = false;
                }

            }, function errorCallback(response) {
            // 请求失败执行代码
                alert("Http Error:TestConfigJson");

            });
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
                $scope.showspider = angular.copy($scope.contactList[$scope.showid]);
                $scope.constantlength = $scope.contactList.length;
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
    