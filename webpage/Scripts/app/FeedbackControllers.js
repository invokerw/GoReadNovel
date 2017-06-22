"use strict";

/* controllers module */

angular.module("BsTableDirective.FeedbackControllers", ["BsTableDirective.Services","ui.bootstrap"])
    .controller("FeedbackCtrl", ["$scope", "$uibModal","BootswatchService","$http", 
        function ($scope, $uibModal, BootswatchService, $http) {
        // scope models
        $scope.page = { Css: "/statics/bootstrap/css/bootstrap.min.css" };
        $scope.contactList = [];
        $scope.progress = { Ready: false };
        // get themes from bootswatch
        BootswatchService.GetAll().success(function (result) {
            $scope.page.Themes = result.themes;
            $scope.page.Css = $scope.page.Themes[15].cssMin;
        });
        $scope.feedbacktype = [
            {name : "书籍问题", type : 0},
            {name : "操作问题", type : 1},
            {name : "其他问题", type : 2},
            {name : "所有问题", type : -1},
        ];
        $scope.solvetype = [
            {name : "未解决问题", type : 0},
            {name : "已解决问题", type : 1},
            {name : "所有问题", type : -1},
        ];
        $scope.select = {
            feedbackselect:$scope.feedbacktype[3],
            solvetypeselect:$scope.solvetype[2],
        };
        $scope.showAllDelBtn = false;
        $scope.$watch('select.feedbackselect',  function(newValue, oldValue) {
            if (newValue === oldValue) 
                return;  
            GenerateData($scope.select.feedbackselect.type,$scope.select.solvetypeselect.type)
        });
        $scope.$watch('select.solvetypeselect',  function(newValue, oldValue) {
            if (newValue === oldValue) 
                return;  
            if (newValue.type == 1) {
                $scope.showAllDelBtn = true;
            }else{
                $scope.showAllDelBtn = false;
            }
            GenerateData($scope.select.feedbackselect.type,$scope.select.solvetypeselect.type)
        });

        // get data for bs-table
        GenerateData($scope.select.feedbackselect.type,$scope.select.solvetypeselect.type)
        // show function for bs-table
        $scope.Show = function (contact) {
            //alert(JSON.stringify(contact));
            OpenNewWindow(contact, true, 0);
        };
        // Solved function for bs-table
        $scope.Solved = function (contact) {
            //alert(JSON.stringify(contact));
            $http({
                method: 'GET',
                url: '/SolvedUserFeedbackJson',
                params:{
                    contactid:contact.feedbackid
                }
            }).then(function successCallback(response) {
                // 请求成功执行代码
                console.log('response.data = ',response.data); 
                if (response.data.code == 1) {
                    GenerateData($scope.select.feedbackselect.type,$scope.select.solvetypeselect.type)
                }   
                // hide progress
            }, function errorCallback(response) {
            // 请求失败执行代码
                alert("Http Get Error:SolvedUserFeedbackJson");

            });
        };
        // remove function for bs-table
        $scope.Remove = function (contact) {
            //alert(JSON.stringify(contact));
            if (confirm("删除该项之后无法恢复，是否确认？")) {  
                $http({
                    method: 'GET',
                    url: '/DelAUserFeedbackJson',
                    params:{
                        contactid:contact.feedbackid
                    }
                }).then(function successCallback(response) {
                    // 请求成功执行代码
                    console.log('response.data = ',response.data); 
                    if (response.data.code == 1) {
                        GenerateData($scope.select.feedbackselect.type,$scope.select.solvetypeselect.type)
                    }  
                    return
                    // hide progress
                }, function errorCallback(response) {
                // 请求失败执行代码
                    alert("Http Get Error:DelAUserFeedbackJson");

                });
            }else {  
                //alert("闲得慌");  
                console.log('闲得慌啊你。。');  
            }  
        };
        $scope.DelAllSolved = function(){
            if (confirm("删除之后无法恢复，是否确认？")) { 
                 $http({
                    method: 'GET',
                    url: '/DelALotUserFeedbackJson',
                    params:{
                        feedbacktype:$scope.select.feedbackselect.type,
                        solved:$scope.select.solvetypeselect.type
                    }
                }).then(function successCallback(response) {
                    // 请求成功执行代码
                    console.log('response.data = ',response.data); 
                    if (response.data.code == 1) {
                        alert("删除成功。")
                        $scope.select.feedbackselect = $scope.feedbacktype[3];
                        $scope.select.solvetypeselect = $scope.solvetype[2];
                        GenerateData($scope.select.feedbackselect.type,$scope.select.solvetypeselect.type)
                    }  
                    // hide progress
                }, function errorCallback(response) {
                // 请求失败执行代码
                    alert("Http Get Error:DelAUserFeedbackJson");

                });
            }else {  
                //alert("闲得慌");  
                console.log('闲得慌啊你。。');  
            } 
        }
        // generate data
        function GenerateData(feedbacktype,solved) {
            $http({
                method: 'GET',
                url: '/GetUserFeedbackJson',
                params:{
                    feedbacktype:feedbacktype,
                    solved:solved
                }
            }).then(function successCallback(response) {
                // 请求成功执行代码
                console.log('response.data = ',response.data);   
                if (response.data.ret == null) {
                    alert("所查数据类型为空。")
                }else{
                    
                    $scope.contactList = response.data.ret;
                }
                // hide progress
                $scope.progress = { Ready: true };
            }, function errorCallback(response) {
            // 请求失败执行代码
                alert("Http Get Error:GetUserFeedbackJson");

            });
        }
        var $ctrl = this;

        function OpenNewWindow(contact,readonly,ty,size){
            //console.log("this.........open");
            var parentElem = undefined;
            //parentSelector ? angular.element($document[0].querySelector('.modal-demo ' + parentSelector)) : undefined;
            var modalInstance = $uibModal.open({
              animation: true,
              ariaLabelledBy: 'modal-title',
              ariaDescribedBy: 'modal-body',
              templateUrl: 'feedbackWindow.html',
              controller: 'ShowFeedBackCtrl',
              controllerAs: '$ctrl',
              size: size,
              appendTo: parentElem,
              resolve: {
                contact:contact,
                opentype:ty,
                readonly:readonly
          
              }
            });
        };
        // init socialite jquery plugin
        Socialite.load();
    }])
    // Please note that $uibModalInstance represents a modal window (instance) dependency.
    // It is not the same as the $uibModal service used above.
    .controller('ShowFeedBackCtrl',["$uibModalInstance","contact", "readonly","$http","opentype",
     function ($uibModalInstance, contact,  readonly,$http,opentype) {
      var $ctrl = this;

      $ctrl.contact = contact;
      //console.log($ctrl.contact);
      $ctrl.readonly = readonly;
      $ctrl.opentype = opentype;
      $ctrl.title = "Feedback";
      //console.log("$ctrl.opentype = ",$ctrl.opentype)
      if ($ctrl.opentype == 0) {
        $ctrl.title = "Feedback Detail";
      }else if($ctrl.opentype == 1){
        $ctrl.title = "Edit Feedback";
      }else if($ctrl.opentype == 2){
        $ctrl.title = "New Feedback";
      }

      $ctrl.ok = function () {
        //readonly = true 是Show
        if($ctrl.opentype == 0) {
            console.log("$ctrl.opentype 0")
        }
        // false 是保存 Edit
        else if($ctrl.opentype == 1) {
            console.log("$ctrl.opentype 1")
        }
        else if($ctrl.opentype == 2) {
            console.log("$ctrl.opentype 2");
        }
        $uibModalInstance.dismiss('cancel');
      };

      $ctrl.cancel = function () {
        $uibModalInstance.dismiss('cancel');
      };
    }]);
    