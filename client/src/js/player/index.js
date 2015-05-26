(function() {
  'use strict';

  var angular = require('angular');

  var ngModule = angular.module('app.player', []);

  ngModule
    .config(function($stateProvider, $urlRouterProvider) {
      $stateProvider
        .state('app.player', {
          url: '/player/:name',
          views: {
            '@': {
              templateUrl: "player/detail.html",
              controller: 'PlayerDetailCtrl'
            }
          }
        });
    })
    .controller('PlayerDetailCtrl', function($scope, $stateParams, Restangular) {
      Restangular.one('players', $stateParams.name).get().then(function(p) {
        if(p.achievements) {
          $scope.achievements = [];
          _.forIn(p.achievements, function(v, k) {
            $scope.achievements.push({
              name: k,
              date: v
            });
          });
          
          p.achievementcount = $scope.achievements.length;
        } else {
          p.achievementcount = 0;
        }

        $scope.player = p;
      });
    });


})();
