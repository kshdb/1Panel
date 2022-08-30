import { CommonModel } from '.';

export namespace Host {
    export interface HostTree {
        label: string;
        children: Array<string>;
    }
    export interface Host extends CommonModel {
        name: string;
        addr: string;
        port: number;
        user: string;
        authMode: string;
        description: string;
    }
    export interface HostOperate {
        id: number;
        name: string;
        addr: string;
        port: number;
        user: string;
        authMode: string;
        privateKey: string;
        password: string;

        description: string;
    }
    export interface ReqSearch {
        info?: string;
    }
}