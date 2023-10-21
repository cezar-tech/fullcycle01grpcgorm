import { Body, Controller, Get, Post } from '@nestjs/common';
import { Observable } from 'rxjs';
import { AppService } from './app.service';
import { ProductDto } from './dto/product.dto';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get('products')
  getProducts(): Observable<
    [{ name: string; description: string; price: number; id: number }]
  > {
    return this.appService.getProducts();
  }

  @Post('products')
  createProduct(@Body() createDto: ProductDto): Observable<object> {
    return this.appService.createProduct(createDto);
  }
}
