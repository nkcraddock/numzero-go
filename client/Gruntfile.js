(function() {
  'use strict';

var cfg = {
  build: {
    path: {
      root: "build/",
      js: "build/js/",
      css: "build/css/",
      fonts: "build/fonts/"
    }
  },
  files: {
    grunt: ['Gruntfile.js'],
    app: {
      js: ['src/js/**/*.js', '!src/js/vendor/**/*.js'],
      css: ['src/js/**/*.css'],
      html: [ 'src/js/**/*.html' ]
    },
    vendor: {
      root: 'src/vendor/bower_components/',
      js: [
        'src/vendor/bower_components/jquery/dist/jquery.js',
        'src/vendor/bower_components/bootstrap/dist/js/bootstrap.js'
      ],
      css: ['src/vendor/bower_components/bootstrap/dist/css/*.min.css'],
      fonts: ['src/vendor/bower_components/bootstrap/dist/fonts/**']
    }
  }
};

module.exports = function(grunt) {

  grunt.initConfig({
    jshint: {
      files: cfg.files.app.js
    },
    watch: {
      build: {
        files: [cfg.files.grunt, 'src/**'],
        tasks: ['build']
      },
      js: {
        files: ['<%= jshint.files %>'],
        tasks: ['jshint']
      }
    },
    clean: [ cfg.build.path.root ],
    copy: {
      build: {
        files: [
        {
          src: cfg.files.vendor.fonts,
          dest: cfg.build.path.fonts,
          flatten: true,
          expand: true
        }]
      },
      vendorjs: {
        files: [{
            src: [ cfg.files.vendor.root + 'modernizr/modernizr.js' ],
            dest: cfg.build.path.js,
            expand: true,
            flatten: true
          }
        ]
      },
      staticContent: {
        files: [{
          src: ['*'],
          dest: cfg.build.path.root,
          cwd: 'src/static/',
          expand: true
        }]
      }
    },
    html2js: {
      main: {
        src: cfg.files.app.html,
        dest: cfg.build.path.js + 'templates.js'
      }
    },
    concat: {
      main: {
        src: cfg.files.app.js,
        dest: cfg.build.path.js + 'app.js'
      },
      maincss: {
        src: cfg.files.app.css,
        dest: cfg.build.path.css + 'app.css'
      },
      vendor: {
        src: cfg.files.vendor.js,
        dest: cfg.build.path.js + 'vendor.js'
      },
      vendorcss: {
        src: cfg.files.vendor.css, 
        dest: cfg.build.path.css + 'vendor.css'
      }
    }
  });

  grunt.loadNpmTasks('grunt-contrib-jshint');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-contrib-copy');
  grunt.loadNpmTasks('grunt-contrib-clean');
  grunt.loadNpmTasks('grunt-contrib-concat');
  grunt.loadNpmTasks('grunt-html2js');


  grunt.registerTask('default', ['build', 'watch:build']);
  grunt.registerTask('build', ['jshint', 'clean', 'html2js', 'concat', 'copy']);


};
})();
