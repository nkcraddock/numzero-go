(function() {
  'use strict';

  var angular = require('angular');

  var ngModule = angular.module('app.home', ['ui.router']);

  ngModule
    .config(function($stateProvider, $urlRouterProvider) {
      $stateProvider
        .state('home', {
          url: '/',
          templateUrl: "home/home.html",
          controller: 'HomeCtrl'
        });
    })
    .controller('HomeCtrl', function($scope) {
      $scope.message = "Shmurda";
    });


})();
