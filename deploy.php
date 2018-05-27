<?php
namespace Deployer;

require 'recipe/common.php';

// Project name
set('application', 'ngobrol');

// Project repository
set('repository', 'git@github.com:marulitua/ngobrol.git');

// [Optional] Allocate tty for git clone. Default value is false.
set('git_tty', true); 

// Shared files/dirs between deploys 
set('shared_files', []);
set('shared_dirs', []);

// Writable dirs by web server 
set('writable_dirs', []);


// Hosts

host('ngobrol.maruli.ga')
    ->set('deploy_path', '~/{{application}}');    
    

// Tasks

desc('Deploy your project');
task('deploy', [
    'local:build',
    'deploy:info',
    'deploy:prepare',
    'deploy:lock',
    'deploy:release',
    'deploy:update_code',
    'local:upload',
    //'deploy:shared',
    //'deploy:writable',
    //'deploy:vendors',
    'deploy:clear_paths',
    'deploy:symlink',
    'remote:restart_service',
    'deploy:unlock',
    'cleanup',
    'success'
]);

// [Optional] If deploy fails automatically unlock.
after('deploy:failed', 'deploy:unlock');

task('local:build', function() {
    runLocally('make build');
    runLocally('make build_assets');
});

task('local:upload', function () {
    upload('public', '{{release_path}}');
    upload('bin', '{{release_path}}');
});

task('remote:restart_service', function () {
    run('sudo systemctl restart ngobrol');
});
