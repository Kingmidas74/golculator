import { AfterViewInit, Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { SplitAreaDirective, SplitComponent } from 'angular-split';
import { Observable } from 'rxjs';
import { AppService } from './app.service';
import { ModalCreateComponent } from './modal-create/modal-create.component';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements AfterViewInit, OnInit {
  title = 'golculator';

  expression = '';

  results:Array<string>;
  operations:Array<any>;
  currentOperation;


  constructor(private appService:AppService,
    private _snackBar: MatSnackBar,
    private dialog: MatDialog) {
    this.results = []
  }

  ngOnInit(): void {
    this.appService.GetOperations().subscribe(result => {
      this.operations = result
    })
  }

  onCodeChange(e) {
    console.log(e)
  }


  ngAfterViewInit(): void {
  /*  ace.config.set("fontSize", "14px");
    ace.config.set('basePath', 'https://unpkg.com/ace-builds@1.4.12/src-noconflict');

    const aceEditor = ace.edit(this.editor.nativeElement);
    aceEditor.session.setValue("");
    aceEditor.setTheme('ace/theme/twilight');
    aceEditor.session.setMode('ace/mode/lua');

    aceEditor.on("change", () => {
      console.log(aceEditor.getValue());
    });*/
  }

  saveAvailable() {
    return ["+","-","*","/"].indexOf(this.currentOperation.Name)<0;
  }

  saveOperation() {
    
    this.appService.SaveOperation(this.currentOperation).subscribe(_=>{
      location.reload();
    });
  }

  calculate() {
    this.appService.CalculateExpression(this.expression)
      .subscribe(
        (result)=> {
          this.results.push(`${this.expression}=${result}`)
        }, 
        err=> {
          this._snackBar.open("Wrong expression",null,{duration:1000});
        })
  }

  openDialog(): void {
    const dialogRef = this.dialog.open(ModalCreateComponent, {
      width: '50rem'
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log('The dialog was closed');
      if(result) {
        this.appService.SaveOperation(result).subscribe(_=>{
          location.reload();
        });
      }
    });
  }

}
