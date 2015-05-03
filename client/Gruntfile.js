(function() {
  'use strict';

module.exports = function(grunt) {

  grunt.initConfig({
    jshint: {
      files: ['Gruntfile.js', 'src/js/**/*.js', '!src/js/vendor/**/*.js' ],
      options: {
        globals: {
          jQuery: true
        }
      }
    },
    watch: {
      build: {
        files: ['src/**'],
        tasks: ['build']
      },
      js: {
        files: ['<%= jshint.files %>'],
        tasks: ['jshint']
      }
    },
    clean: [ 'build/' ],
    copy: {
      build: {
        files: [
          {
            src: ['**'],
            dest: 'build/',
            cwd: 'src',
            expand: true
          }
        ]
      }
    }
  });

  grunt.loadNpmTasks('grunt-contrib-jshint');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-contrib-copy');
  grunt.loadNpmTasks('grunt-contrib-clean');


  grunt.registerTask('default', ['build', 'watch:build']);
  grunt.registerTask('build', ['jshint', 'clean', 'copy:build']);

};
})();
