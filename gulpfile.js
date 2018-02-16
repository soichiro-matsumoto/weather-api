var gulp        = require('gulp'),
    plumber     = require('gulp-plumber'),
    uglify      = require("gulp-uglify"),
    browserify  = require('browserify'),
    babelify    = require('babelify'),
    source      = require('vinyl-source-stream');
 
// トランスパイル
gulp.task('browserify', function() {
    browserify('./resources/js/jsx/app.jsx', { debug: true  })
        .transform(babelify, { presets: ['es2015', 'react'] })
        .bundle()
        .on("error", function (err) { console.log("Error : " + err.message); })
        .pipe(source('bundle.js'))
        .pipe(gulp.dest('./public/js/'))
        .on('end',function(){
            gulp.src(["./public/js/bundle.js"])
                .pipe(plumber())
                .pipe(uglify({mangle: false}))
                .pipe(gulp.dest("./public/js/min"))
        })
});
 
// watch 
gulp.task('watch',function(){
    gulp.watch('./resources/js/jsx/*.jsx', ['browserify']);
});
 
gulp.task("default",['watch']);