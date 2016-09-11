angular.module('gatorloopWebApp')

.service('dashboardService', function($http) {
    var api = "/api/";
    return {
        sendStopSignal: function() {
            return $http.get(api + "stop");
        },
        get: function(url) {
            return $http.get(api + url);
        }
    }
})
