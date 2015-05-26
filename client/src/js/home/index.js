(function() {
  'use strict';

  var angular = require('angular');

  var ngModule = angular.module('app.home', []);

  ngModule
    .config(function($stateProvider, $urlRouterProvider) {
      $stateProvider
        .state('app.home', {
          url: '/',
          views: {
            '@': {
              templateUrl: "home/home.html",
              controller: 'HomeCtrl'
            }
          }
        })
        .state('app.dashboard', {
          url: '/leaders',
          views: {
            '@': {
              templateUrl: 'home/dashboard.html',
              controller: 'DashboardCtrl'
            }
          }
        });
    })
    .controller('HomeCtrl', function($scope) {

    })
    .controller('DashboardCtrl', function($scope, $state, Restangular) {
      $scope.gotoplayer = function(player) {
        $state.go('app.dashboard', { "name": player.name });
      };
      Restangular.all('players').getList().then(function(players) {
        var i = 0;
        $scope.players = _.chain(players)
          .sortBy(function(p) { return p.score * -1; })
          .take(5)
          .map(function(p) {
            if(p.achievements) {
              p.achievementcount = Object.keys(p.achievements).length;
            } else {
              p.achievementcount = 0;
            }
            p.rank = i++;
            return p;
          })
          .value();
      });
    });


})();
