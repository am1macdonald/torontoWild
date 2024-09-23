// application router

const router = ((): Router => {
    const routes: Record<string, (...args: any) => void> = {}

    const addRoute = (path: string, handler: (...args: any) => void) => {
        routes[path] = handler
    }

    const navigateTo = (path: string) => {
        history.pushState({}, '', path)
        routes[path]()
    }

    const init = () => {
        document.body.addEventListener('click', (e) => {
            if (e.target && e.target instanceof HTMLAnchorElement) {
                e.preventDefault()
                navigateTo(e.target.href)
            }
        })

        window.addEventListener('popstate', () => {
            routes[window.location.pathname]()
        })
    }
    return {
        addRoute,
        navigateTo,
        init
    }
})()


export default router

export interface Router {
    addRoute: (path: string, handler: (...args: any) => void) => void
    navigateTo: (path: string) => void
    init: () => void
}
