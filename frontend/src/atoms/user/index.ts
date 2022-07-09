import { atom, useAtomValue, useSetAtom } from 'jotai'
export interface UserInfo {
    token: string // jwt token (sub in gsi)
    name: string // user first name
    email: string // user email
    picture: string // default picture provided by google
}

export const BLANK_USER: UserInfo = {
    token: '',
    name: '',
    email: '',
    picture: ''
}

export const tokenAtom = atom<string>('')
export const nameAtom = atom<string>('')
export const emailAtom = atom<string>('')
export const pictureAtom = atom<string>('')

export const userInfoAtom = atom<UserInfo, UserInfo>(
    get => ({
        token: get(tokenAtom),
        name: get(nameAtom),
        email: get(emailAtom),
        picture: get(pictureAtom)
    }),
    (get, set, payload) => {
        set(tokenAtom, payload.token)
        set(nameAtom, payload.name)
        set(emailAtom, payload.email)
        set(pictureAtom, payload.picture)
    }
)

export const useGetUserInfo = () => {
    return useAtomValue(userInfoAtom)
}

export const useUserActions = () => {
    const setUserInfo = useSetAtom(userInfoAtom)
    return {
        setUserInfo: (payload: UserInfo) => setUserInfo(payload)
    }

}