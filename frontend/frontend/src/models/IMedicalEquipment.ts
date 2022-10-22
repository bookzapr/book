import { TypeofUseInterface } from "./ITypeofUse";


export interface MedicalEquimentInterface {

    ID: number;
    Name:string;

    TypeofUseID: number;
    TypeofUse: TypeofUseInterface;
   
}