import { NestFactory } from '@nestjs/core';
import { SwaggerModule, DocumentBuilder } from '@nestjs/swagger';
import { AppModule } from './app.module';
import { ConfigService } from '@nestjs/config';
import {
  BadRequestException,
  ValidationError,
  ValidationPipe,
} from '@nestjs/common';
import * as bodyParser from 'body-parser';
import { ERROR_CODES } from './contants/error-codes';
import { CommonResponse } from './common/dto/common-response.dto';
import { HttpLoggerInterceptor } from './http-logger.interceptor';
import { NestExpressApplication } from '@nestjs/platform-express';

async function bootstrap() {
  const app = await NestFactory.create<NestExpressApplication>(AppModule);
  const configService = app.get(ConfigService);

  app.set('trust proxy', true);

  app.enableCors({
    origin: '*',
    methods: '*',
    credentials: true,
  });

  app.setGlobalPrefix('/api');
  app.enableShutdownHooks(['SIGINT', 'SIGTERM']);

  app.use(bodyParser.json({ limit: '10mb' }));
  app.use(bodyParser.urlencoded({ limit: '10mb', extended: true }));

  app.useGlobalPipes(
    new ValidationPipe({
      whitelist: true,
      forbidNonWhitelisted: true,
      transform: true,
      exceptionFactory: (errors: ValidationError[]) => {
        const details = errors.map((error) => ({
          field: error.property,
          message: Object.values(error.constraints || {}).join(', '),
        }));

        const errorCode = ERROR_CODES.VALIDATION_FAILED;

        return new BadRequestException(
          new CommonResponse(
            undefined,
            errorCode.message,
            errorCode.code,
            details,
          ),
        );
      },
    }),
  );

  app.useGlobalInterceptors(new HttpLoggerInterceptor());

  const config = new DocumentBuilder()
    .setTitle('Socketing API Documentation')
    .setDescription(
      'This documentation provides details about the Socketing service.',
    )
    .setVersion('1.0')
    .addBearerAuth()
    .build();
  const documentFactory = () => SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('api/docs', app, documentFactory);

  const port = configService.get<number>('PORT');
  await app.listen(port);

  if (process.send) {
    process.send('ready');
    console.log('Worker process is ready.');
  }

  console.log(`Application is running on port ${port}`);
}
bootstrap();
