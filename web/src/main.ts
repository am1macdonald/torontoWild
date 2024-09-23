import './style.css'
import router, {Router} from "./lib/router.ts";
import home from './routes'
import login from './routes/login'

const appConfig: {
    router: Router,
} = {
    router: router
}



function main() {
    initRouter();
    appConfig.router.navigateTo(window.location.pathname);
}

function initRouter() {
    appConfig.router.addRoute('/', home)
    appConfig.router.addRoute('/login', login)
    appConfig.router.init();
}

main();
