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
        $scope.search = {
            searchType:[
                {name : "ID Name or Author", type : 0},
                {name : "Novel Type", type : 1},
            ],
            key:"",
            novelType:[
            "玄幻小说","仙侠小说","都市小说",
            "言情小说","网游小说","科幻小说",
            "历史小说","灵异小说","其他小说",
            ],
            selectedSearchType :"",
            showType:false
        };
        $scope.info = {
            novelCount:0,
            nBegin:0,
            nNum:100,
            nowpage:0,
            allpage:1,
            NextClickShow:true,
            PreClickShow:false
        };
        $scope.$watch('search.selectedSearchType',  function(newValue, oldValue) {
            if (newValue === oldValue) return;  
            console.log("search.selectedSearchType new value:",newValue);
            if (newValue.type == 1) {
                $scope.search.showType = true;
            }else{
                $scope.search.showType = false;
            }
        });
        $scope.$watch('search.key',  function(newValue, oldValue) {
            if (newValue === oldValue) return;  
            if (newValue == "") return;
            console.log("key new value:",newValue);
            $http({
                    method: 'GET',
                    url: '/GetUltimateSearchNovelsJson',
                    params: {
                        type:$scope.search.selectedSearchType.type,
                        key:$scope.search.key
                    },
                }).then(function successCallback(response) {
                    // 请求成功执行代码
                    if(response.data.code == 1){
                        //删除成功
                        console.log('response.data = ',response.data);
                        $scope.contactList = response.data.ret;
                        $scope.info.novelCount = $scope.contactList.length;
                        $scope.info.nowpage = 0;
                        $scope.info.allpage = 1;
                        $scope.info.NextClickShow = false;
                        $scope.info.PreClickShow = false;
                    } 
                    else
                    {
                        console.log("没有搜索结果");
                    }
                 }, function errorCallback(response) {
                    // 请求失败执行代码
                    alert("Https Get Error:GetUltimateSearchNovelsJson");

                });
        });
        //console.log(" $scope.info.novelCount", $scope.info.novelCount)
        // get data for bs-table
        GenerateData($scope.info.nBegin,$scope.info.nNum);
        //获取小说数量
        $scope.NextPageData = function(){
            if ($scope.info.nowpage >= $scope.info.allpage - 1) return;
            $scope.progress = { Ready: false };
            $scope.info.NextClickShow = false;
            $scope.info.nowpage++;
            console.log("$scope.info = ",$scope.info);
            var begin = $scope.info.nBegin + $scope.info.nowpage*$scope.info.nNum;
            if ($scope.info.allpage - 1 == $scope.info.nowpage) //翻到了最后一页
            {
                GenerateData(begin, $scope.info.novelCount - begin);
                $scope.info.NextClickShow = false;
            }
            else  //没有到最后一页
            {
                GenerateData(begin, $scope.info.nNum);
                $scope.info.NextClickShow = true;
            }
            $scope.info.PreClickShow = true;
        }
        $scope.PrePageData = function(){
            if ($scope.info.nowpage <= 0) return;
            $scope.progress = { Ready: false };
            $scope.info.PreClickShow = false;
            $scope.info.nowpage--;
            var begin = $scope.info.nBegin + $scope.info.nowpage*$scope.info.nNum;
            if (0 == $scope.info.nowpage) //翻到了第一页
            {
                $scope.info.PreClickShow = false;
            }
            GenerateData(begin, $scope.info.nNum);
            if ($scope.info.nowpage >= 1) //翻到了第一页
            {
                $scope.info.PreClickShow = true;
            }
            $scope.info.NextClickShow = true;
        }
        GetNovelsCountService.GetAll().success(function (result) {
            $scope.info.novelCount  = result.ret;
            //向上取整
            $scope.info.allpage = Math.ceil($scope.info.novelCount/$scope.info.nNum);
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