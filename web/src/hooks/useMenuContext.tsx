import { createContext, useContext, useState, type ReactElement } from "react"

type Pages = "button" | "stats"

interface MenuContextInterface {
    page: Pages
    setPage: (p: Pages) => void
}

export const MenuContext = createContext<MenuContextInterface | undefined>(undefined)

export function useMenuContext() {
    const context = useContext(MenuContext)
    if (context === undefined) {
        throw new Error("useMenuContext must be used within a MenuProvider");
    }
    return context
}

export function MenuProvider({ children }: { children: ReactElement }) {
    const [page, setPage] = useState<Pages>("button")

    return <MenuContext.Provider value={{ page: page, setPage: setPage }}>
        {children}
        <Menu />
    </MenuContext.Provider>
}

function Menu() {
    const menu = useMenuContext()
    const onUndo = () => {

    }

    const onLogout = () => {
        document.location.reload()
    }

    const onPageSwitch = () => {
        switch (menu.page) {
            case "button":
                menu.setPage("stats")
                break
            case "stats":
                menu.setPage("button")
                break
        }
    }

    const pageSwitchLabel = menu.page === "button" ? "calendar" : "button"

    return (<div className="menu">
        <button className='button-base' onClick={onUndo}>undo</button>
        <button className='button-base' onClick={onPageSwitch}>{pageSwitchLabel}</button>
        <button className='button-base' onClick={onLogout}>logout</button>
    </div>)
}