'use strict';

angular.module('gatorloopWebApp')
    .controller('dashboardCTRL', function($scope, dashboardService) {

        $scope.newVelocity = 0.00;

      $scope.sendStopSignal = function() {
          dashboardService.sendStopSignal().success(function(data) {
              alert(data);
          }).error(function(data) {
              alert("Error", data);
          });
      };

      $scope.getCurrentVelocity = function() {
          dashboardService.getData("velocity").success(function(data) {
              alert(data);
          }).error(function(data) {
              alert("Error", data);
          });
      };

      $scope.saveCurrentVelocity = function() {
          console.log($scope.newVelocity);
          var v = JSON.stringify({velocity: $scope.newVelocity});
          dashboardService.postData("velocity", v).success(function(data) {
            alert(data);
          }).error(function(data) {
             alert("Error", data);
          });
      }

      $scope.getCurrentPressure = function() {
          dashboardService.getCurrentPressure().success(function(data) {
              alert(data);
          }).error(function(data) {
              alert("Error", data);
          });
      };

      $scope.getCurrentPosition = function() {
          dashboardService.getCurrentPosition().success(function(data) {
              alert(data);
          }).error(function(data) {
              alert("Error", data);
          });
      };

      $scope.getCurrentAcceleration = function() {
          dashboardService.getCurrentAcceleration().success(function(data) {
              alert(data);
          }).error(function(data) {
              alert("Error", data);
          });
      };

      $scope.getCurrentTemperature = function() {
          dashboardService.getCurrentTemperature().success(function(data) {
              alert(data);
          }).error(function(data) {
              alert("Error", data);
          });
      };

    });
