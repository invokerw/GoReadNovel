"use strict";

/* controllers module */

angular.module("BsTableDirective.Controllers", ["BsTableDirective.Services","ui.bootstrap"])
    .controller("ExampleCtrl", ["$scope", "BootswatchService", function ($scope, BootswatchService) {
        // scope models
        $scope.page = { Css: "/statics/bootstrap/css/bootstrap.min.css" };
        $scope.contactList = [];
        $scope.progress = { Ready: false };
        // get themes from bootswatch
        BootswatchService.GetAll().success(function (result) {
            // set themes to scope
            $scope.page.Themes = result.themes;
            // hide progress
            $scope.progress = { Ready: true };
        });
        // get data for bs-table
        $scope.contactList = GenerateData(20);
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
        function GenerateData(count) {
            var dataList = [];
            for (var i = 0; i < count; i++)
            {
                // create object
                var item = {
                    FirstName: "FirstName " + i,
                    LastName: "LastName " + i,
                    BornDate: "2014-07-10"
                };
                // add to list
                dataList.push(item);
            }
            return dataList;
        }
        // init socialite jquery plugin
        Socialite.load();
    }])
    
    .controller("NovelInfoCtrl", ["$scope", "$uibModal","BootswatchService", "$http", "GetNovelsCountService",
        function ($scope,$uibModal, BootswatchService, $http, GetNovelsCountService) {
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
        $scope.info = {
            novelCount:0,
            nBegin:0,
            nNum:100,
            nowpage:1,
            allpage:1
        };
        function GetData(){
        }
        //console.log(" $scope.info.novelCount", $scope.info.novelCount)
        // get data for bs-table
        GenerateData($scope.info.nBegin,$scope.info.nNum);
        //获取小说数量

        GetNovelsCountService.GetAll().success(function (result) {
            $scope.info.novelCount  = result.ret;
        });

        // show function for bs-table
        $scope.Show = function (contact) {
            //alert(JSON.stringify(contact));
            OpenNewWindow(contact,true,0);
        };
        // edit function for bs-table
        $scope.Edit = function (contact) {
            //alert(JSON.stringify(contact));
            OpenNewWindow(contact,false,1);
        };
        // remove function for bs-table
        $scope.Remove = function (contact) {
            //alert(JSON.stringify(contact));
            if (confirm("你确定删除该项？")) {  
                //alert("点击了确定");  
                $http({
                    method: 'GET',
                    url: '/GetDeleteNovleID',
                    params: {
                        novelid:contact.id
                    },
                }).then(function successCallback(response) {
                    // 请求成功执行代码
                    if(response.data.code == 1){
                        //删除成功
                        console.log('response.data = ',response.data);  
                        GenerateData(0,100);
                    } 
                 }, function errorCallback(response) {
                    // 请求失败执行代码
                    alert("Https Get Error:GetDeleteNovleID");

                });
            }  
            else {  
                //alert("闲得慌");  
                console.log('闲得慌啊你。。');  
             }  
        };
        //打开小窗口呀
        var $ctrl = this;
        $ctrl.items = ['item1', 'item2', 'item3'];
        $scope.selected = ""
        $scope.Open = function(){
            OpenNewWindow(undefined, false, 2);
        }

        function OpenNewWindow(contact,readonly,ty,size){
            //console.log("this.........open");
            var parentElem = undefined;
            //parentSelector ? angular.element($document[0].querySelector('.modal-demo ' + parentSelector)) : undefined;
            var modalInstance = $uibModal.open({
              animation: true,
              ariaLabelledBy: 'modal-title',
              ariaDescribedBy: 'modal-body',
              templateUrl: 'myModalContent.html',
              controller: 'ModalInstanceCtrl',
              controllerAs: '$ctrl',
              size: size,
              appendTo: parentElem,
              resolve: {
                items: function () {
                    return $ctrl.items;
                },
                contact:contact,
                opentype:ty,
                readonly:readonly
          
              }
            });

            modalInstance.result.then(function (selectedItem) {
              $scope.selected = selectedItem;
            }, function () {
              console.log('Modal dismissed at: ' + new Date());
            });
        };
        
        // 获取小说table信息
        function GenerateData(begin,count) {
            var dataList = [];
            $http({
                method: 'GET',
                url: '/GetNovelTableInfoJson',
                params: {
                    begin:begin,
                    num:count
                },
            }).then(function successCallback(response) {
                // 请求成功执行代码
                dataList = response.data.ret; //这里需要修改
                //console.log('response.data = ',response.data);   
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
    }])
    
    // Please note that $uibModalInstance represents a modal window (instance) dependency.
    // It is not the same as the $uibModal service used above.
    .controller('ModalInstanceCtrl',["$uibModalInstance","contact", "items","readonly","$http","opentype",
     function ($uibModalInstance, contact, items, readonly,$http,opentype) {
      var $ctrl = this;
      $ctrl.items = items;
      $ctrl.contact = contact;
      $ctrl.readonly = readonly;
      $ctrl.opentype = opentype;
      $ctrl.title = "Novel";
      console.log("$ctrl.opentype = ",$ctrl.opentype)
      if ($ctrl.opentype == 0) {
        $ctrl.title = "Show Novel";
      }else if($ctrl.opentype == 1){
        $ctrl.title = "Edit Novel";
      }else if($ctrl.opentype == 2){
        $ctrl.title = "New Novel";
      }
      $ctrl.selected = {
        item: $ctrl.items[0]
      };

      $ctrl.ok = function () {
        //readonly = true 是Show
        if($ctrl.opentype == 0) {
            $uibModalInstance.close($ctrl.selected.item);
        }
        // false 是保存 Edit
        else if($ctrl.opentype == 1) {
            console.log("$ctrl.readonly false")
            $http({
                method: 'GET',
                url: '/GetEditNovelJson',
                params: {novel:$ctrl.contact},
            }).then(function successCallback(response) {
                // 请求成功执行代码

                if (response.data.code == 1) {
                    console.log("edit ok",response.data);
                    $uibModalInstance.close($ctrl.selected.item);
                }
                else {
                     console.log("edit err",response.data);
                     alert(response.data);
                }

            }, function errorCallback(response) {
            // 请求失败执行代码
                alert("Http Get Error:GetTNovelTableInfo");
            });

        }
        else if($ctrl.opentype == 2) {
            console.log("new novel");
            $http({
                method: 'GET',
                url: '/GetANewNovelJson',
                params: {novel:$ctrl.contact},
            }).then(function successCallback(response) {
                // 请求成功执行代码

                if (response.data.code == 1) {
                    console.log("new novel ok",response.data);
                    $uibModalInstance.close($ctrl.selected.item);
                }
                else {
                     console.log("new novel err",response.data);
                     alert(response.data);
                }

            }, function errorCallback(response) {
            // 请求失败执行代码
                alert("Http Get Error:GetTNovelTableInfo");
            });
        }
      };

      $ctrl.cancel = function () {
        $uibModalInstance.dismiss('cancel');
      };
    }]);