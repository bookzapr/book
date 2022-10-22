import { DoctorInterface } from "./IDoctor";
import { WorkplaceInterface } from "./IWorkplace";
import { MedicalEquimentInterface } from "./IMedicalEquipment";

export interface BorrowInterface {

    ID: number;
    // entity
    DoctorID : number;
    Doctor: DoctorInterface;

    WorkplaceID : number;
    Workplace: WorkplaceInterface;

    MedicalEquipmentID : number;
    MedicalEquipment: MedicalEquimentInterface;

    Quant: number;

    BorrowDate: Date;
    ReturnDate: Date;

   //แสดงข้อมูลมาแสดงมาจาก หลังบ้าน
   }