angular.module('gatorloopWebApp')

.service('dashboardService', function($http) {
    var api = "/api/";
    return {
        killPod: function() {
            return $http.get(api + "killpower");
        },
        emergencyStop: function() {
          return $http.get(api + "stop");
        },
        get: function(url) {
            return $http.get(api + url);
        }
    }
})
