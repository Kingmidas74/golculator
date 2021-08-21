import { Component, ViewChild } from '@angular/core';
import { SplitAreaDirective, SplitComponent } from 'angular-split';
import { AppService } from './app.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'golculator';

  expression = '';

  constructor(private appService:AppService) {

  }

  calculate() {
    this.appService.CalculateExpression(this.expression)
      .subscribe(console.log, console.error)
  }

}
