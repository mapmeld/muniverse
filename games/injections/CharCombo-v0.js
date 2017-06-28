(function() {

  // Game depends on a certain framerate.
  faketime.animationFrameRate = 60;

  var c2 = window.construct2api;

  window.muniverse = {
    init: function(options) {
      return pollAndWait(100, function() {
        return window.game;
      }).then(function() {
        faketime.pause();
        faketime.advance(100);

        // Used to detect game over.
        ['best0', 'best1'].forEach((k) => c2.globalVar(k).setValue(-1));
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      var done = ['best0', 'best1'].some((k) => c2.globalVar(k).getValue() >= 0);
      return Promise.resolve(done);
    },
    score: function() {
      return Promise.resolve(c2.globalVar('score').getValue());
    }
  };

})();
