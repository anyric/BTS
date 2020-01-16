var gulp = require("gulp")
var shell = require("gulp-shell")

// compile new binary with source change
gulp.task("install-binary", shell.task([
    'go install github.com/anyric/bts'
]))

// restarts the server on new binary change
gulp.task("restart-supervisor", shell.task([
    'supervisorctl restart btsserver'
]))

gulp.task('watch', function() {
    // watch the code base for all changes
    gulp.watch("*", gulp.series('install-binary', 'restart-supervisor'));
});

// setup default task
gulp.task('default', gulp.series('watch'));
