export interface Group {
    id: number
    name: string
    links: Links[]
}

export interface Links {
    id: number
    title: string
    url: string
    group_id: number
}