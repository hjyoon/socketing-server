import { IsString, IsOptional, IsArray, IsInt } from 'class-validator';
import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';
import { swaggerExampleUrls } from 'src/config/swagger-example-urls';

export class CreateEventRequestDto {
  @ApiProperty({
    description: 'Title of the event',
    example: 'Music Festival',
  })
  @IsString()
  title: string;

  @ApiPropertyOptional({
    description: 'Thumbnail URL for the event',
    example: swaggerExampleUrls.eventThumbnail,
  })
  @IsOptional()
  @IsString()
  thumbnail?: string;

  @ApiProperty({
    description: 'Location or venue of the event',
    example: 'Central Park',
  })
  @IsString()
  place: string;

  @ApiProperty({
    description: 'Main cast or performers of the event',
    example: 'Famous Band',
  })
  @IsString()
  cast: string;

  @ApiPropertyOptional({
    description: 'Age limit for the event (if any)',
    example: 18,
  })
  @IsOptional()
  @IsInt()
  ageLimit?: number;

  @ApiPropertyOptional({
    description: 'SVG data for the event design or layout',
    example:
      '<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100"><circle cx="50" cy="50" r="40" stroke="black" stroke-width="3" fill="red" /></svg>',
  })
  @IsOptional()
  @IsString()
  svg?: string;

  @ApiPropertyOptional({
    description: 'The starting time for ticketing in ISO 8601 format',
    type: String,
    example: '2024-11-20T10:00:00.000Z',
  })
  @IsOptional()
  @IsString()
  ticketingStartTime?: Date;

  @ApiPropertyOptional({
    description: 'List of event dates in ISO 8601 format',
    type: [String], // Swagger에서는 Date를 String으로 표시
    example: ['2024-12-01T19:00:00.000Z', '2024-12-02T19:00:00.000Z'],
  })
  @IsArray()
  @IsOptional()
  eventDates?: Date[];
}
