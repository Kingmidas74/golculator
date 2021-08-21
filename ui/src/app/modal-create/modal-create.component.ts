import { Component, OnInit } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-modal-create',
  templateUrl: './modal-create.component.html',
  styleUrls: ['./modal-create.component.scss']
})
export class ModalCreateComponent {

  operation = {
    Name:"",
    ArgumentsCount:0,
    Code:""
  }

  constructor(
    public dialogRef: MatDialogRef<ModalCreateComponent>
    ) {}

  onNoClick(): void {
    this.dialogRef.close();
  }

  onChange(e) {
    this.operation.Code = e;
  }
  
  Save() {
    this.operation.ArgumentsCount = parseInt(this.operation.ArgumentsCount.toString())    
    this.dialogRef.close(this.operation);
  }

}
