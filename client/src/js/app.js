
(function() {
  'use strict';

  var angular = require('angular');

  var app = angular.module('app', [
    'ui.router',
    'templates-main',
    'app.home',
    'app.layout'
  ]);

  app.config(function($locationProvider) {
    $locationProvider.html5Mode({
      enabled: true,
      requireBase: false
    });
  });

  app.config(function($urlRouterProvider) {
    $urlRouterProvider.otherwise('/');
  });


})();
